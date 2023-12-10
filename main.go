package main

import (
	"gin_init/global"
	"gin_init/internal/routers"
	"gin_init/internal/setup"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func init() {
	setup.ApplicationStartInit()
}
func main() {
	global.Logger.Infof("%s: go_gin_init/%s", "@Nuitke", "go_gin_init")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
