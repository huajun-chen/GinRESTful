package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRouter 用户路由
func UserRouter(r *gin.RouterGroup) {
	UserRouter := r.Group("/user")
	{
		UserRouter.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}
