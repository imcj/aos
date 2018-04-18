package logs

import (
	"time"

	"aos/pkg/setting"
	"aos/pkg/tool"

	"github.com/apex/log"

	"github.com/gin-gonic/gin"
)

func Logger(logger *log.Entry) gin.HandlerFunc {

	return func(c *gin.Context) {
		// fmt.Println("xxxx:", c.GetHeader("Connection"))
		tool.NewUniqueIDAsync()
		uuid := tool.GetUID()

		// fmt.Println(" X-Response-ID:", uuid)
		setting.Logger = logger.WithField("X-Response-ID", uuid)

		t := time.Now()
		c.Next()
		latency := time.Since(t)
		setting.Logger.Infof("接口："+c.HandlerName()+"，请求时间为：", latency)
	}
}
