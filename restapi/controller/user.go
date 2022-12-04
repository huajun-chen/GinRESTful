package controller

import (
	"GinRESTful/restapi/dao"
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 用户登录
func Login(c *gin.Context) {
	loginForm := forms.LoginForm{}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		// 参数异常处理
		utils.HandleValidatorError(c, err)
		return
	}
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Msg:  "success",
	})
}

func GetUserList(c *gin.Context) {
	// 获取参数
	userListForm := forms.UserListForm{}
	if err := c.ShouldBindQuery(&userListForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	page, pageSize := utils.PageZero(userListForm.Page, userListForm.PageSize)
	total, userList, err := dao.GetUserListDao(page, pageSize)
	if err != nil {
		response.Response(c, response.ResponseStruct{
			Code: global.SelectDBErrCode,
			Msg:  global.SelectDBErr,
		})
		return
	}
	// 获取数据为空
	if total == 0 {
		response.Response(c, response.ResponseStruct{
			Code: global.DataEmptyCode,
			Msg:  global.DataEmpty,
		})
		return
	}
	// 获取数据正常
	data := make(map[string]interface{})
	data["total"] = total
	data["values"] = userList
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Data: data,
	})
}
