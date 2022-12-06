package utils

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

// HandleValidatorError 处理字段校验异常
func HandleValidatorError(c *gin.Context, err error) {
	//如何返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		response.Response(c, response.ResponseStruct{
			Code: global.ParameterErrCode,
			Msg:  global.ParameterErr,
			Data: err.Error(),
		})
		return
	}
	data := removeTopStruct(errs.Translate(global.Trans))
	response.Response(c, response.ResponseStruct{
		Code: global.ParameterErrCode,
		Msg:  global.ParameterErr,
		Data: data,
	})
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for filed, err := range fileds {
		// 从文本的圆点(.)开始切分   处理后"mobile": "mobile为必填字段"  处理前: "PasswordLoginForm.mobile": "mobile为必填字段"
		rsp[filed[strings.Index(filed, ".")+1:]] = err
	}
	return rsp
}
