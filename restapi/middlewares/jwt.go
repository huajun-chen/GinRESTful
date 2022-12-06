package middlewares

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	TokenExpired     = errors.New("Token is expired")            // 令牌已过期
	TokenNotValidYet = errors.New("Token not active yet")        // 令牌尚未激活
	TokenMalformed   = errors.New("That's not even a token")     // 这不是一个令牌
	TokenInvalid     = errors.New("Couldn't handle this token:") // 无法处理此令牌
)

// CustomClaims 自定义声明
type CustomClaims struct {
	ID          uint   // ID
	Name        string // 名字
	AuthorityID uint   // 权限
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

// JWTAuth JWT认证
// 参数：
//		无
// 返回值：
//		gin.HandlerFunc：Gin的处理程序
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息
		// 这里前端需要把token存储到cookie或者本地localSstorage中
		// 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.Response(c, response.ResponseStruct{
				Code: global.NoTokenCode,
				Msg:  global.NoToken,
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// 解析Token信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.Response(c, response.ResponseStruct{
					Code: global.AuthExpiredCode,
					Msg:  global.AuthExpired,
				})
				c.Abort()
				return
			}
			response.Response(c, response.ResponseStruct{
				Code: global.InvalidTokenCode,
				Msg:  global.InvalidToken,
			})
			c.Abort()
			return
		}
		// Gin的上下文记录claims和userId的值
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		c.Next()
	}
}

// NewJWT NEW JEW
// 参数：
//		无
// 返回值：
//		*JEW：jwt的指针
func NewJWT() *JWT {
	return &JWT{[]byte(global.Settings.JWTKey.SigningKey)}
}

// CreateToken 创建一个token
// 参数：
//		claims：声明
// 返回值：
//		string：Token字符串
//		error：错误信息
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token
// 参数：
//		tokenString：Token字符串
// 返回值：
//		CustomClaims：自定义声明
//		error：错误信息
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// RefreshToken 更新Token
// 参数：
//		tokenString：Token字符串
// 返回值：
//		string：更新后的Token字符串
//		error：错误信息
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
