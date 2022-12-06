package utils

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/middlewares"
	"GinRESTful/restapi/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// CreateToken 生成Token信息
// 参数：
//		c：gin.Context的指针
//		id：用户ID
//		role：角色的值
//		name：用户名
// 返回值：
//		string：Token字符串
func CreateToken(c *gin.Context, id uint, role int, name string) string {
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID:          id,
		Name:        name,
		AuthorityID: uint(role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// Token过期时间
			ExpiresAt: time.Now().Unix() + int64(global.Settings.JWTKey.TokenExpiration),
			Issuer:    global.Settings.Name,
		},
	}
	// 生成Token
	token, err := j.CreateToken(claims)
	if err != nil {
		response.Response(c, response.ResponseStruct{
			Code: global.CreateTokenFailCode,
			Msg:  global.CreateTokenFail,
		})
		return ""
	}
	return token
}
