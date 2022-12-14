package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResStruct 响应结构体
type ResStruct struct {
	// omitempty(省略)：字段如果没有值就不显示此字段
	Code int         `json:"code,omitempty"` // 自定义响应状态码
	Msg  string      `json:"msg,omitempty"`  // 响应信息
	Data interface{} `json:"data,omitempty"` // 响应数据
}

// Response 统一封装响应函数
// 参数：
//		c：gin.Context的指针
//		response：相应的结构体
// 返回值：
//		无
func Response(c *gin.Context, response ResStruct) {
	// 所有请求的响应系统状态码都返回200，在code字段自定义状态码
	c.JSON(http.StatusOK, response)
	return
}
