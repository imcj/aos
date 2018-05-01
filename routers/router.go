package routers

import (
	"aos/bindService"
	"aos/middleware/logs"
	"aos/middleware/panicHandle"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "aos/docs"

	"aos/pkg/setting"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://xxx.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	setting.Logger = setting.GrayLog()

	r.Use(logs.Logger(setting.Logger))

	r.Use(gin.Recovery())
	r.Use(panicHandle.CatchError())

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	testApi := container.ProjectController
	// := new(controller.TestApi)

	c1 := r.Group("/v1")
	c1.GET("/secret/:access_key", testApi.GetS)
	c1.GET("/dbtest", testApi.GetDbTest)
	c1.GET("/servicetest", testApi.GetServiceTest)
	return r
}
