package main

import (
	"fileshared/src/conf"
	"fileshared/src/handler"
	"fmt"
	"net/http"
	"os"
)

const (
	ExitCodeLoadConfError    = 1
	ExitCodeStartServerError = 2
)

func main() {
	if !conf.InitConf() {
		fmt.Printf("load conf failed\n")
		os.Exit(ExitCodeLoadConfError)
	}
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", conf.GlobalConf.Port),
		Handler: &handler.ServerHandler{},
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("start server failed, error:%s", err.Error())
		os.Exit(ExitCodeStartServerError)
	}
}
