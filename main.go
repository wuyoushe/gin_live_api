package main

import (
	"fmt"
	"net/http"

	"github.com/wuyoushe/gin_live_api/pkg/setting"
	"github.com/wuyoushe/gin_live_api/routers"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
