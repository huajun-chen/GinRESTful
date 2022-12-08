package middlewares

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// I18n i18n中间件
// 参数：
//		无
// 返回值：
//		gin.HandlerFunc：Gin的处理程序
func I18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的语言信息，示例：zh-CN
		acceptLanguage := c.Request.Header.Get("Accept-Language")
		// 判断语言信息是否为空，判断语言信息是否在全部语言中，避免获取到恶意参数，注意!取反
		if acceptLanguage == "" || !strings.Contains(global.Settings.Language.AllLanguage, acceptLanguage) {
			// 使用系统默认的语言
			acceptLanguage = global.Settings.Language.LanguageType
		}
		// 文件完整路径
		filePath := fmt.Sprintf("%s/%s.json", global.Settings.Language.Tranfilepath, acceptLanguage)
		// 根据Accept-Language读取对应的json文件
		translations, err := utils.ReadJSON(filePath)
		if err != nil {
			response.Response(c, response.ResponseStruct{
				Code: 10006,
				Msg:  global.I18nMap["10006"],
			})
			c.Abort()
			return
		}
		// 传递全局变量
		global.I18nMap = translations

		c.Next()
	}
}
