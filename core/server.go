package core

import (
	"ginProject/global"
	"ginProject/initialize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := initialize.Routers()
	address := ":8085"
	s := initServer(address, Router)
	global.GVA_LOG.Info("服务正在启动，port", zap.String("adress", address))
	global.GVA_LOG.Info(s.ListenAndServe().Error())
	//global.GVA_LOG.Error(s.ListenAndServe().Error())
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	//time.Sleep(10 * time.Microsecond)
}
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
