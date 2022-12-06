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
		userRouterToken.PUT("/info/:id", controller.ModifyUserInfo) // 修改用户信息
		userRouterToken.GET("/info/:id", controller.GetMyselfInfo)  // 用户查看自己的信息
		// 需要Token和权限的接口
		{
			userRouterTokenAdmin := userRouterToken.Group("")
			userRouterTokenAdmin.Use(middlewares.IsAdminAuth())
			userRouterTokenAdmin.GET("/list", controller.GetUserList)    // 用户列表
			userRouterTokenAdmin.DELETE("/info/:id", controller.DelUser) // 删除用户信息
		}
	}
}
