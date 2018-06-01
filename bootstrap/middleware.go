package bootstrap

import (
	"github.com/gin-gonic/gin"
)

type MiddlewareCommand struct {

}

func (c MiddlewareCommand)Execute() {
	middlewares["CROSHTTPMiddleware"] = middleware.CROSHTTPMiddleware()
	// RegisterEngine(func (engine *gin.Engine) {
		
	// })
}