package bootstrap

import (
	"fmt"
	"time"
	// "fmt"
	"log"
	"syscall"

	"github.com/aos-stack/aos/bootstrap/middleware"
	"github.com/aos-stack/env"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Registry func(engine *gin.Engine)

type HTTPServerCommand struct{}

func (c HTTPServerCommand) Execute() {
	router := gin.New()

	var middlewares []string = make([]string, 0)
	for _, m := range viper.Get("http.middleware").([]interface{}) {
		middlewares = append(middlewares, m.(string))
	}

	middleware.UseGinHTTPMiddlewares(middlewares, router)

	if env.Get() > 0 {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	Registry(router)
	initService(router)
}

func initService(router *gin.Engine) {

	endless.DefaultReadTimeOut = viper.GetDuration("http.server.timeout.read") * time.Second
	endless.DefaultWriteTimeOut = viper.GetDuration("http.server.timeout.write") * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", viper.GetInt("http.server.port"))

	log.Printf("http.server.port %d", viper.GetInt("http.server.port"))
	server := endless.NewServer(endPoint, router)

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	log.Printf("over success")
}
