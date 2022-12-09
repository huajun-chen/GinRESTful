package controller

import (
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetSystemInfo 获取系统信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func GetSystemInfo(c *gin.Context) {
	// CPU
	cpuStruct, err := utils.CPUInfo()
	if err != nil {
		response.Response(c, response.ResponseStruct{
			Code: 10023,
			Msg:  global.I18nMap["10023"],
		})
		return
	}
	// 内存
	memStruct, err := utils.MemInfo()
	if err != nil {
		response.Response(c, response.ResponseStruct{
			Code: 10024,
			Msg:  global.I18nMap["10024"],
		})
		return
	}
	// 硬盘
	diskStruct, err := utils.DiskInfo()
	if err != nil {
		response.Response(c, response.ResponseStruct{
			Code: 10025,
			Msg:  global.I18nMap["10025"],
		})
		return
	}

	data := forms.System{
		CPU:    cpuStruct,
		Memory: memStruct,
		Disk:   diskStruct,
	}
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Data: data,
	})
}
