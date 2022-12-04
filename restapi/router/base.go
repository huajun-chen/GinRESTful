package router

import (
	"GinRESTful/restapi/controller"
	"github.com/gin-gonic/gin"
)

// BaseRouter 基础路由
func BaseRouter(r *gin.RouterGroup) {
	baseRouter := r.Group("/base")
	{
		// 无需Token的接口
		baseRouter.GET("/captcha", controller.GetCaptcha) // 验证码
	}
}
