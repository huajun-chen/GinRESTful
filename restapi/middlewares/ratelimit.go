package middlewares

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// RateLimit 限流中间件，限制并发的请求数量
// 参数：
//		无
// 返回值：
//		gin.HandlerFunc：Gin的处理程序
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		rateLimtFloat64 := float64(global.Settings.RateLimit)
		rateLimtInt64 := int64(global.Settings.RateLimit)
		// 创建限流器
		limiter := ratelimit.NewBucketWithRate(rateLimtFloat64, rateLimtInt64)
		if limiter.TakeAvailable(1) == 0 {
			response.Response(c, response.ResponseStruct{
				Code: 10022,
				Msg:  global.I18nMap["10022"],
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
