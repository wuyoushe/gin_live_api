package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuyoushe/gin_live_api/pkg/logging"
	"github.com/wuyoushe/gin_live_api/pkg/setting"
	"github.com/wuyoushe/gin_live_api/routers"
)

func main() {
	engine := gin.Default()
	engine.Use(logging.LoggerToFile())
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
