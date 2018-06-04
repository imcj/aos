package middleware

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/aos-stack/aos/bootstrap/middleware"
	"testing"
	"github.com/stretchr/testify/mock"
	// ".."
)

type EngineMock struct {
	mock.Mock
}

func TestRegisterGinHTTPMiddleware(t *testing.T) {
	middleware.RegisterGinHTTPMiddleware("X-REQUEST-ID", func (c *gin.Context)gin.HandlerFunc {
		return func (c *gin.Context) {

		}
	})
	requestID := middleware.GetGinHTTPMiddleware("X-REQUEST-ID")
}

func T(t *testing.T) {
	// middleware.UseMiddlewares({"X-REQUEST-ID"}, engine)
}