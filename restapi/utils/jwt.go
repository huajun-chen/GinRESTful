package utils

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

// CreateToken 生成Token信息
// 参数：
//		c：gin.Context的指针
//		id：用户ID
//		role：角色的值
//		name：用户名
// 返回值：
//		string：Token字符串
func CreateToken(c *gin.Context, id uint, role int, name string) string {
	j := NewJWT()
	claims := CustomClaims{
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

// IsInBlacklist 判断Token是否在黑名单中（true：在，false不在）
// 参数：
//		tokenStr：Token字符串
// 返回值：
//		bool：Token是否在黑名单里，true：在，false不在
func IsInBlacklist(tokenStr string) bool {
	tokenMD5 := MD5(tokenStr)
	value, _ := RedisGetStr(tokenMD5)
	// 如果在Redis中通过key获取的值为空，说明此Token未加入黑名单
	if value == "" {
		return false
	}
	return true
}
