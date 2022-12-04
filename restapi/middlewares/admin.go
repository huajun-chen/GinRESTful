package middlewares

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"github.com/gin-gonic/gin"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Token信息
		claims, _ := c.Get("claims")
		// 获取当前用户信息
		currentUser := claims.(*CustomClaims)
		// 判断是否具有权限
		if currentUser.AuthorityID != 1 {
			response.Response(c, response.ResponseStruct{
				Code: global.AuthInsufficientCode,
				Msg:  global.AuthInsufficient,
			})
			// 中断之后的中间件
			c.Abort()
			return
		}
		// 继续向下执行
		c.Next()
	}
}
