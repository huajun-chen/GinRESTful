package router

import (
	"GinRESTful/restapi/controller"
	"GinRESTful/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

// Routerv1 V1版本路由
func Routerv1(r *gin.RouterGroup) {
	// 基础功能的路由
	baseRouter := r.Group("/base")
	{
		// 无需Token的接口
		baseRouter.GET("/captcha", controller.GetCaptcha) // 验证码
		// 需要Token的接口
	}

	// 用户模块路由
	userRouter := r.Group("/user")
	{
		// 无需Token的接口
		userRouter.POST("/register", controller.Register) // 注册
		userRouter.POST("/login", controller.Login)       // 登录
		// 需要Token的接口
		userRouterToken := userRouter.Group("")
		userRouterToken.Use(middlewares.JWTAuth())
		{
			userRouterToken.GET("/list", middlewares.IsAdminAuth(), controller.GetUserList) // 用户列表
		}
	}
}
