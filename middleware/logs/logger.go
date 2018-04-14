package logs

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		fmt.Println(latency)
		// setting.Logger.Info("接口：" + c.HandlerName() + "，请求时间为：" + strconv.FormatInt(int64(latency)/1000000, 10) + "ms")
	}
}
