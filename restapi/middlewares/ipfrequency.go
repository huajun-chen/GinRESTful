package middlewares

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"github.com/gin-gonic/gin"
	"time"
)

// Frequency 限制某个IP访问的频率
// 参数：
//		无
// 返回值：
//		gin.HandlerFunc：Gin的处理程序
func Frequency() gin.HandlerFunc {
	counter := make(map[string]int)
	expiration := make(map[string]time.Time)
	return func(c *gin.Context) {
		// 获取客户端IP地址
		ip := c.ClientIP()
		// 获取当前时间
		nowTime := time.Now()
		// 如果客户端IP地址的计数器已存在，则检测是否到期
		if count, ok := counter[ip]; ok {
			// 计数已到期，重置计数器
			if expiration[ip].Before(nowTime) {
				counter[ip] = 1
				// 默认限制1分钟时长
				expiration[ip] = nowTime.Add(time.Duration(global.Settings.UserInfo.TimeLimit) * time.Minute)
			} else if count > global.Settings.UserInfo.IpFrequency {
				// 计数器未到期，但已超过最大频率次数
				response.Response(c, response.ResponseStruct{
					Code: 10021,
					Msg:  global.I18nMap["10021"],
				})
				c.Abort()
				return
			} else {
				// 计数器未到期，增加计数器
				counter[ip]++
			}
		} else {
			// 客户端IP地址的计数器不存在，创建计数器
			counter[ip] = 1
			expiration[ip] = nowTime.Add(time.Duration(global.Settings.UserInfo.TimeLimit) * time.Minute)
		}
		c.Next()
	}
}
