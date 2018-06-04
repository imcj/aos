package middleware

import (
	// "time"
	"github.com/gin-gonic/gin"
)

func RequestChainHTTPMiddleware()gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("X-REQUEST-ID", "ID")
		c.Next()
	}
}