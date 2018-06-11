package middleware

import (
	"fmt"
	// "github.com/apex/log"
	"github.com/gin-gonic/gin"
)

var ginMiddlewares map[string]gin.HandlerFunc

func RegisterGinHTTPMiddleware(name string, middleware gin.HandlerFunc) {
	if nil == ginMiddlewares {
		ginMiddlewares = make(map[string]gin.HandlerFunc)
	}
	ginMiddlewares[name] = middleware
}

func GetGinHTTPMiddleware(name string) gin.HandlerFunc {
	return ginMiddlewares[name]
}

func UseGinHTTPMiddlewares(middlewares []string, engine *gin.Engine) {
	// log.Debug("Will ")
	fmt.Println("will http Middlewares")
	for _, middleware := range middlewares {
		fmt.Println("Use http middleware " + middleware)
		engine.Use(ginMiddlewares[middleware])
	}
}
