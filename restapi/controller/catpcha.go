package controller

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

// base64Captcha  缓存对象
var store = base64Captcha.DefaultMemStore

// GetCaptcha 获取验证码
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	// bs64是图片的base64编码
	id, bs64, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10007"], err.Error())
		response.Response(c, response.ResponseStruct{
			Code: 10007,
			Msg:  global.I18nMap["10007"],
		})
		return
	}
	data := make(map[string]interface{})
	data["captcha_id"] = id
	data["captcha_path"] = bs64
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Data: data,
	})
}
