package controller

import (
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 用户登录
func Login(c *gin.Context) {
	loginForm := forms.LoginForm{}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		// 参数异常处理
		utils.HandleValidatorError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
