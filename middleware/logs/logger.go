package logs

import (
	"time"

	"github.com/aos-stack/aos/pkg/setting"
	"github.com/aos-stack/aos/pkg/tool"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		tool.NewUniqueIDAsync()
		uuid := tool.GetUID()
		logger := setting.GrayLog()
		logger = logger.WithField("X-Response-ID", uuid)
		c.Set("logger", logger)
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		logger.Infof("interface url:"+c.Request.URL.Path+"，times:", latency)
	}
}
