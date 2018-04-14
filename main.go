package main

import (
	"aos/routers"
	"aos/secret"
	"fmt"
	"log"
	"net/http"
	"syscall"

	_ "aos/docs"

	"aos/pkg/setting"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type ResponseObject struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// TODO: 处理HTTP响应包括错误的公共方法
func Dump(c *gin.Context, err error, object interface{}) {
	responseObject := ResponseObject{
		1,
		err.Error(),
		object,
	}
	if nil != err {
		c.JSON(http.StatusOK, responseObject)
	} else {
		c.JSON(http.StatusOK, responseObject)
	}
}

func CreateSecretFromRequest(c *gin.Context) secret.Secret {
	accessKey := c.PostForm("access_key")
	if accessKey == "" {
		accessKey = c.Param("access_key")
	}
	accessSecret := c.DefaultQuery("access_secret", "")

	return secret.Secret{
		accessKey,
		accessSecret,
	}
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

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService 127.0.0.1:6001

// @license.name MIT
// @license.url 127.0.0.1:6001

// @BasePath /v1
func main() {

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		fmt.Println("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	// // setting.Logger.Info("I am tester shengji")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// // TODO: 对象依赖配置放到专门的模块
	// var (
	// 	secretDAO           = persistence.NewSecretDAO(client)
	// 	secretServiceFacade = secret.NewSecretServiceFacadeImpl(
	// 		secretDAO,
	// 		secret.NewSecretFactory(),
	// 	)
	// )

	// router := gin.Default()

	// gin.SetMode(setting.RunMode)

	// // TODO: Controller 放置到专门的模块内
	// router.POST("/secret", func(c *gin.Context) {
	// 	authentication := CreateSecretFromRequest(c)

	// 	newSecret, err := secretServiceFacade.Add(authentication)
	// 	if nil != err {
	// 		fmt.Println(err)
	// 	}

	// 	c.JSON(http.StatusOK, newSecret)
	// })

	// router.Use(ResponseMiddleware())
	// router.Use(logs.Logger())

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// apiv1 := router.Group("/v1")

	// apiv1.GET("/secret/:access_key", getS)

	// router.Run(":6001")
}
