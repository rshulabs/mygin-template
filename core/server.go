package core

/*
实现思路
1.创建gin实例，绑定路由组
2.http库 开启监听并实现优雅退出
*/

import (
	"backup/global"
	"backup/router"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	r := setRouter()
	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
	log.Println("server exited")
}
func setRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	router.SetApiGroupRoutes(apiGroup)
	return r
}
