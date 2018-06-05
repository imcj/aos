package bootstrap

import (
	"fmt"
	// "fmt"
	"log"
	"syscall"

	"github.com/aos-stack/aos/bootstrap/middleware"

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

	Registry(router)

	fmt.Println("==========")
	fmt.Println(viper.GetDuration("http.server.timeout.read"))
	fmt.Println("==========")

	endless.DefaultReadTimeOut = viper.GetDuration("http.server.timeout.read")
	endless.DefaultWriteTimeOut = viper.GetDuration("http.server.timeout.write")
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", viper.GetInt("http.server.port"))

	log.Printf("http.server.port %d", endPoint)
	server := endless.NewServer(endPoint, router)

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	log.Printf("123")
}
