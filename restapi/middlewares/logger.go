package middlewares

import (
	"GinRESTful/restapi/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// GinLogger 接收gin框架默认的日志
// 参数：
//		无
// 返回值：
//		gin.HandlerFunc：Gin的处理程序
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		zap.L().Info("",
			zap.Int("status", c.Writer.Status()),                                 // 请求状态
			zap.String("method", c.Request.Method),                               // 请求方式
			zap.String("path", c.Request.URL.Path),                               // 请求路径
			zap.String("query", c.Request.URL.RawQuery),                          // 请求参数
			zap.String("ip", c.ClientIP()),                                       // 请求IP
			zap.String("user-agent", c.Request.UserAgent()),                      // 用户代理
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()), // 错误信息
			zap.Duration("cost", time.Since(start)),                              // 处理时间
		)
	}
}

// GinRecovery recover项目可能出现的panic，并使用zap记录相关日志
// 参数：
//		无
// 返回值：
//		gin.HandlerFunc：Gin的处理程序
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.Lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					return
				}
				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("stack", string(debug.Stack())),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
