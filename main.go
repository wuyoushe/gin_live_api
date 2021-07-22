package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuyoushe/gin_live_api/models"
	"github.com/wuyoushe/gin_live_api/pkg/logging"
	"github.com/wuyoushe/gin_live_api/pkg/setting"
	"github.com/wuyoushe/gin_live_api/routers"
)

// func init() {
// 	setting.Setup()
// 	models.Setup()
// 	logging.Setup()

// }

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

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

// func main() {
// 	endless.DefaultReadTimeOut = setting.ReadTimeout
// 	endless.DefaultWriteTimeOut = setting.WriteTimeout
// 	endless.DefaultMaxHeaderBytes = 1 << 20
// 	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

// 	server := endless.NewServer(endPoint, routers.InitRouter())
// 	server.BeforeBegin = func(add string) {
// 		log.Printf("Actual pid is %d", syscall.Getpid())
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Printf("Server err: %v", err)
// 	}
// }
