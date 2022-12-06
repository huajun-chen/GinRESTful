package initialize

import (
	"GinRESTful/restapi/global"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// InitTrans validator信息翻译
// 参数：
//		locale：语言环境（中文/英文/...）
// 返回值：
//		err：错误信息
func InitTrans(locale string) (err error) {
	// 修改Gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		// 中文翻译
		zhT := zh.New()
		// 英文翻译
		enT := en.New()
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			_ = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}
