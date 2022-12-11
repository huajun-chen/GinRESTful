package service

import (
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// base64Captcha  缓存对象
var store = base64Captcha.DefaultMemStore

// SerGetCaptcha 业务层：获取验证码
// 参数：
//		无
// 返回值：
//		response.ResStruct：响应的结构体
func SerGetCaptcha() response.ResStruct {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	// bs64是图片的base64编码
	id, bs64, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10007"], err)
		failStruct := response.ResStruct{
			Code: 1007,
			Msg:  global.I18nMap["10007"],
		}
		return failStruct
	}
	data := forms.CaptchaReturn{
		CaptchaId:   id,
		CaptchaPath: bs64,
	}
	succStruct := response.ResStruct{
		Code: 200,
		Data: data,
	}
	return succStruct
}
