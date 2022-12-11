package v1

import (
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/service"
	"github.com/gin-gonic/gin"
)

// ConGetCaptcha 控制层：获取验证码
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetCaptcha(c *gin.Context) {
	JSONStr := service.SerGetCaptcha()
	response.Response(c, JSONStr)
}
