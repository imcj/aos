package routers

import (
	"github.com/aos-stack/aos/bindService"
	"github.com/aos-stack/aos/middleware/logs"
	"github.com/aos-stack/aos/middleware/panicHandle"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/aos-stack/aos/docs"

	"github.com/aos-stack/aos/pkg/setting"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(logs.Logger())
	r.Use(gin.Recovery())
	r.Use(panicHandle.CatchError())

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	containerInstance := container.GetContainer()
	c1 := r.Group("/v1")

	//test
	c1.GET("/graylog", containerInstance.TestApi.TestGraylog)
	c1.GET("/name", containerInstance.TestApi.TestName)

	//static
	r.Static("/public", "./public")
	return r
}
