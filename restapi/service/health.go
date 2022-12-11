package service

import (
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/utils"
	"go.uber.org/zap"
	"net/http"
)

// SerGetSystemInfo 业务层：获取系统信息
// 参数：
//		无
// 返回值：
//		response.ResStruct：响应的结构体
func SerGetSystemInfo() response.ResStruct {
	// CPU
	cpuStruct, err := utils.CPUInfo()
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10023"], err)
		failStruct := response.ResStruct{
			Code: 10023,
			Msg:  global.I18nMap["10023"],
		}
		return failStruct
	}
	// 内存
	memStruct, err := utils.MemInfo()
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10024"], err)
		failStruct := response.ResStruct{
			Code: 10024,
			Msg:  global.I18nMap["10024"],
		}
		return failStruct
	}
	// 硬盘
	diskStruct, err := utils.DiskInfo()
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10025"], err)
		failStruct := response.ResStruct{
			Code: 10025,
			Msg:  global.I18nMap["10025"],
		}
		return failStruct
	}

	data := forms.SystemReturn{
		CPU:    cpuStruct,
		Memory: memStruct,
		Disk:   diskStruct,
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Data: data,
	}
	return succStruct
}
