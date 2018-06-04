package middleware

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSGinHTTPMiddleware()gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://xxx.com"
		// },
		MaxAge: 12 * time.Hour,
	})
}