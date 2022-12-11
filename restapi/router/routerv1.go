package router

import (
	v1 "GinRESTful/restapi/controller/v1"
	"GinRESTful/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

// Routerv1 V1版本路由
// 参数：
//		r：Gin的路由分组的指针
// 返回值：
//		无
func Routerv1(r *gin.RouterGroup) {
	// 基础功能的路由
	baseRouter := r.Group("/base")
	{
		// 无需Token的接口
		//baseRouter.GET("/captcha", controller.GetCaptcha) // 验证码
		baseRouter.GET("/captcha", v1.ConGetCaptcha) // 验证码
		// 需要Token的接口
		baseRouterToken := baseRouter.Group("")
		baseRouterToken.Use(middlewares.JWTAuth())
		//baseRouterToken.GET("/health", controller.GetSystemInfo) // 系统资源使用情况，CPU，内存，硬盘
		baseRouterToken.GET("/health", v1.ConGetSystemInfo) // 系统资源使用情况，CPU，内存，硬盘
	}

	// 用户模块路由
	userRouter := r.Group("/user")
	{
		// 无需Token的接口
		//userRouter.POST("/register", controller.Register) // 注册
		userRouter.POST("/register", v1.ConRegister) // 注册
		//userRouter.POST("/login", controller.Login)  // 登录
		userRouter.POST("/login", v1.ConLogin) // 登录
		// 需要Token的接口
		userRouterToken := userRouter.Group("")
		userRouterToken.Use(middlewares.JWTAuth())
		//userRouterToken.PUT("/info/:id", controller.ModifyUserInfo) // 修改用户信息
		userRouterToken.PUT("/info/:id", v1.ConModifyUserInfo) // 修改用户信息
		//userRouterToken.GET("/info/:id", controller.GetMyselfInfo)  // 用户查看自己的信息
		userRouterToken.GET("/info/:id", v1.ConGetMyselfInfo) // 用户查看自己的信息
		//userRouterToken.DELETE("/logout", controller.Logout)        // 登出
		userRouterToken.DELETE("/logout", v1.ConLogout) // 登出
		// 需要Token和权限的接口
		{
			userRouterTokenAdmin := userRouterToken.Group("")
			userRouterTokenAdmin.Use(middlewares.IsAdminAuth())
			//userRouterTokenAdmin.GET("/list", controller.GetUserList)    // 用户列表
			userRouterTokenAdmin.GET("/list", v1.ConGetUserList) // 用户列表
			//userRouterTokenAdmin.DELETE("/info/:id", controller.DelUser) // 删除用户信息
			userRouterTokenAdmin.DELETE("/info/:id", v1.ConDelUser) // 删除用户信息
		}
	}
}
