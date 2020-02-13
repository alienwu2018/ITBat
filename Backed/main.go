package main

import (
	"ITBat/pkg/setting"
	"ITBat/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.SetMode(setting.RUNMODE)
	routersInit := router.InitRoute()
	readTimeout := setting.READTIMEOUT
	writeTimeout := setting.WRITETIMEOUT
	endPoint := fmt.Sprintf(":%d", setting.HTTPPORT)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	log.Fatal(server.ListenAndServe())
}
