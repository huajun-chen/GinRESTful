package router

import (
	"GinRESTful/restapi/controller"
	"GinRESTful/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

// UserRouter 用户路由
func UserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/user")
	{
		// 无需Token的接口
		userRouter.POST("/login", controller.Login) // 登录

		// 需要Token的接口
		userRouterToken := userRouter.Group("")
		userRouterToken.Use(middlewares.JWTAuth())
		{
			userRouterToken.GET("/list", controller.GetUserList) // 用户列表
		}
	}
}
