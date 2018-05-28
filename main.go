package main

import (
	"fmt"
	"log"
	"syscall"

	_ "sparta.gaodun.com/docs"
	"sparta.gaodun.com/routers"

	"sparta.gaodun.com/pkg/setting"

	"os"

	"sparta.gaodun.com/pkg/utils"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ResponseObject struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// TODO:
// 输出HTTP处理日志
// 配置权限、用户状态等对象容器
// 输出RequestID等处理调用链路
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func preSigUsr1() {
	fmt.Println("endless.PRE_SIGNAL ....")
}
func postSigUsr1() {
	fmt.Println("endless.POST_SIGNAL ... ")
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService 127.0.0.1:6001

// @license.name MIT
// @license.url 127.0.0.1:6001

// @BasePath /v1
func main() {

	env := os.Getenv("SYSTEM_ENV")
	if env != "" {
		viper.Set("env", env)
	}

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	// init log
	setting.LoadConfig()

	// init db
	if err := utils.InitEngine(); err != nil {
		fmt.Println("数据库连接异常：", err)
		// panic(err)
		os.Exit(0)
	}

	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	handle := routers.InitRouter()

	server := endless.NewServer(endPoint, handle)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
