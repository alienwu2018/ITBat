package main

import (
	"ITBat/pkg/logging"
	"ITBat/pkg/setting"
	"ITBat/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(setting.ServerCfg.RUNMODE)
	routersInit := router.InitRoute()
	readTimeout := setting.ServerCfg.READTIMEOUT
	writeTimeout := setting.ServerCfg.WRITETIMEOUT
	endPoint := fmt.Sprintf(":%d", setting.ServerCfg.HTTPPORT)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	logging.Logger.Printf("[info] start http server listening %s", endPoint)
	logging.Logger.Fatal(server.ListenAndServe())
}
