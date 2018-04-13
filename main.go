package main

import (
	"aos/persistence"
	"aos/secret"
	"fmt"
	"net/http"
	"os"

	_ "aos/docs"

	"github.com/apex/log"
	"github.com/apex/log/handlers/graylog"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	e, _ := graylog.New("udp://g02.graylog.gaodunwangxiao.com:5504")

	t := text.New(os.Stderr)

	log.SetHandler(multi.New(e, t))

	ctx := log.WithFields(log.Fields{
		"item":    "ginlog",
		"message": "test",
		"ahhh":    "xxxx",
	})
	ctx.Info("upload")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// TODO: 对象依赖配置放到专门的模块
	var (
		secretDAO           = persistence.NewSecretDAO(client)
		secretServiceFacade = secret.NewSecretServiceFacadeImpl(
			secretDAO,
			secret.NewSecretFactory(),
		)
	)

	router := gin.Default()

	// gin.SetMode()

	// TODO: Controller 放置到专门的模块内
	router.POST("/secret", func(c *gin.Context) {
		authentication := CreateSecretFromRequest(c)

		newSecret, err := secretServiceFacade.Add(authentication)
		if nil != err {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, newSecret)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(ResponseMiddleware())

	apiv1 := router.Group("/v1")

	apiv1.GET("/secret/:access_key", getS)

	router.Run(":6001")
}

// @Summary 获取S
// @Produce  json
// @Param access_key path string true "秘钥KEY"
// @Success 200 {string} json "{"status": 1,"message": "","result": {"access_key": "xxx","access_secret": ""}}"
// @Router /secret/{access_key} [get]
func getS(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// TODO: 对象依赖配置放到专门的模块
	var (
		secretDAO           = persistence.NewSecretDAO(client)
		secretServiceFacade = secret.NewSecretServiceFacadeImpl(
			secretDAO,
			secret.NewSecretFactory(),
		)
	)

	fmt.Println("GET /secret/:access_key")

	authentication := CreateSecretFromRequest(c)
	fmt.Println("Access key is " + authentication.AccessKey)
	fmt.Println("Access secret is " + authentication.AccessSecret)
	_, err := secretServiceFacade.Authenticate(authentication)
	if nil != err {
		fmt.Println(err)
	}
	c.JSON(200, ResponseObject{
		1,
		"",
		authentication,
	})
}
