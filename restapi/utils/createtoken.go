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
