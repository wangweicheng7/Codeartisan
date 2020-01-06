package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/wangweicheng7/Codeartisan/pkg/setting"
	"github.com/wangweicheng7/Codeartisan/router"
)

func init() {
	setting.Setup()
}

func main() {
	fmt.Println("hello world!")

	router := router.Init()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
