package routers

import (
	"aos/middleware/panicHandle"

	"github.com/gin-gonic/gin"

	_ "aos/docs"
	"aos/middleware/logs"

	"aos/pkg/setting"

	"aos/controller"

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

	c1 := r.Group("/v1")
	c1.GET("/secret/:access_key", controller.GetS)

	return r
}
