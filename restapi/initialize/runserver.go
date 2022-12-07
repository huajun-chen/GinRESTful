package initialize

import (
	"GinRESTful/restapi/global"
	"context"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunServer 运行服务，实现优雅关机
// 参数：
//		无
// 返回值：
//		无
func RunServer() {
	router := InitRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.Settings.Port),
		Handler: router,
	}
	go func() {
		// 开启一个goroutine连接服务
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	color.Red("Shutdown Server...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
	color.Red("Server exiting...")
}
