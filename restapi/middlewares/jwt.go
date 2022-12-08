package middlewares

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuth JWT认证
// 参数：
//		无
// 返回值：
//		gin.HandlerFunc：Gin的处理程序
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwt鉴权取头部信息Authorization登录时回返回token信息
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			response.Response(c, response.ResponseStruct{
				Code: 10009,
				Msg:  global.I18nMap["10009"],
			})
			c.Abort()
			return
		}
		// 按空格分隔Authorization内容（Bearer token信息）
		bearerToken := strings.Split(authorization, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			response.Response(c, response.ResponseStruct{
				Code: 10011,
				Msg:  global.I18nMap["10011"],
			})
			c.Abort()
			return
		}

		j := utils.NewJWT()
		// 解析Token信息
		claims, err := j.ParseToken(bearerToken[1])
		if err != nil {
			if err == utils.TokenExpired {
				response.Response(c, response.ResponseStruct{
					Code: 10010,
					Msg:  global.I18nMap["10010"],
				})
				c.Abort()
				return
			}
			response.Response(c, response.ResponseStruct{
				Code: 10011,
				Msg:  global.I18nMap["10011"],
			})
			c.Abort()
			return
		}
		// 判断Token是否在黑名单中（true：在，false不在）
		ok := utils.IsInBlacklist(bearerToken[1])
		if ok {
			response.Response(c, response.ResponseStruct{
				Code: 10011,
				Msg:  global.I18nMap["10011"], // Token在黑名单中，定义为失效
			})
			c.Abort()
			return
		}

		// Gin的上下文记录claims
		c.Set("claims", claims)
		// 用户ID
		c.Set("userId", claims.ID)
		// 用户Token
		c.Set("token", bearerToken[1])
		// Token到期时间戳
		c.Set("tokenExpiresAt", claims.ExpiresAt)
		c.Next()
	}
}
