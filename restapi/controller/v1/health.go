package v1

import (
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/service"
	"github.com/gin-gonic/gin"
)

// ConGetSystemInfo 控制层：获取系统信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetSystemInfo(c *gin.Context) {
	JSONStr := service.SerGetSystemInfo()
	response.Response(c, JSONStr)
}
