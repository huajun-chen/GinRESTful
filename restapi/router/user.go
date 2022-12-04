package router

import (
	"GinRESTful/restapi/controller"
	"github.com/gin-gonic/gin"
)

// UserRouter 用户路由
func UserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/login", controller.Login)
		userRouter.GET("/list", controller.GetUserList)
	}
}
