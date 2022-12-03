package main

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/initialize"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化YAML配置
	initialize.InitConfig()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", global.Settings.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
