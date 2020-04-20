package main

import (
	"github.com/XuJinTao1996/coredns-management/pkg/setting"
	"github.com/XuJinTao1996/coredns-management/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           router,
		ReadHeaderTimeout: setting.ReadTimeout,
		WriteTimeout:      setting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	s.ListenAndServe()
}
