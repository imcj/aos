package bootstrap

import (
	// "github.com/spf13/viper"
	// "fmt"
	// "github.com/aos-stack/env"
	"github.com/apex/log"
	_ "github.com/spf13/viper/remote"
	"github.com/aos-stack/aos/bootstrap/middleware"
	middlewares "github.com/aos-stack/aos/middleware"

)

// type EnvCommand struct {}

// func (c EnvCommand)Execute() {
// 	struct 
// 	env.SetProvider(provider)
// }

type RemoteConfigCommand struct {}

func (r RemoteConfigCommand)Execute() {
	
}

type GinHTTPMiddlewareCommand struct {
}

func (c GinHTTPMiddlewareCommand)Execute() {
	log.Info("Did load GinHTTPMiddlewareCommand")
	middleware.RegisterGinHTTPMiddleware("CORSGinHTTPMiddleware", middlewares.CORSGinHTTPMiddleware())
	middleware.RegisterGinHTTPMiddleware("RequestChainHTTPMiddleware", middlewares.RequestChainHTTPMiddleware())
}