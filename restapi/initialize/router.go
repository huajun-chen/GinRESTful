package initialize

import (
	"GinRESTful/restapi/middlewares"
	"GinRESTful/restapi/router"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
// 参数：
//		*gin.Engine：Gin引擎的指针
// 返回值：
//		无
func InitRouter() *gin.Engine {
	Router := gin.Default()
	// 跨域中间件
	Router.Use(middlewares.Cors())
	// 注册zap相关中间件
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	// 注册i18n国际化中间件
	Router.Use(middlewares.I18n())
	// 路由分组
	APIGroup := Router.Group("/api")
	{
		// v1版本路由
		APIv1 := APIGroup.Group("/v1")
		{
			router.Routerv1(APIv1)
		}
	}
	return Router
}
