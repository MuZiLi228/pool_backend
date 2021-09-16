package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"pool_backend/bootstrap"
	"pool_backend/configs"
	_ "pool_backend/docs"
	"pool_backend/src/global"
	"pool_backend/src/routers"
	"pool_backend/src/ws"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	_ "github.com/joho/godotenv/autoload"
)

var (
	r = gin.Default()
)

// @title swagger 接口文档
// @version 2.0
// @description
// @contact.name
// @contact.url
// @contact.email
// @license.name MIT
// @license.url https://www.baidu.com
// @host 192.168.1.163:7001
// @BasePath
func main() {
	//系统初始化
	bootstrap.Init()

	defer global.DB.DbRClose()
	defer global.DB.DbWClose()

	// 初始化 socketio
	server := socketio.NewServer(nil)
	go func() {
		if err := server.Serve(); err != nil {
			global.Logger.Error("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	// 初始化 http 服务
	router := routers.InitRouter()

	new(ws.SocketHandler).RegisterSocket(router, server)

	router.StaticFS("/public", http.Dir("./asset"))

	srv := &http.Server{
		Addr:    ":" + configs.ProjectPort(),
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("Server Start Error: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error("Server Shutdown Error:", err)
	}
	global.Logger.Debug("Server Shutdown")

}
