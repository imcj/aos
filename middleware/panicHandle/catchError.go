package panicHandle

import (
	"aos/errors"
	"aos/pkg/setting"

	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type HTTPError interface {
	HTTPStatus() int
}

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stackInfo := "堆栈信息：" + string(debug.Stack())
				setting.GrayLog(map[string]interface{}{"stackInfo": stackInfo}).Infof("", stackInfo)
				c.JSON(200, errors.New(errors.SYSERR, errors.GetInfo()[errors.SYSERR]))

				c.Abort()
			}
		}()
		c.Next()
	}
}
