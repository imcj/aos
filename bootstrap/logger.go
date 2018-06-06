package bootstrap

import (
	"fmt"
	"os"
	"strings"

	aosContainer "github.com/aos-stack/aos/container"
	"github.com/apex/log"
	"github.com/apex/log/handlers/graylog"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Logger *log.Entry
var grayLog log.Handler

type hl string

func (*hl) HandleLog(e *log.Entry) error {
	return nil
}

var l hl = "nil"

type LoggerCommand struct {
}

func (c LoggerCommand) Execute() {
	log.Info("Did load LoggerCommand")

	grayLog = getGrayLog()
	Logger = GrayLog()
	aosContainer.ContainerSet("logger", Logger)

}

func GrayLog(newFields ...map[string]interface{}) *log.Entry {

	isUseGray := viper.IsSet("logger.handlers.graylog")
	isShowConsole := viper.IsSet("logger.handlers.console")
	t := text.New(os.Stderr)
	if !isUseGray {
		log.SetHandler(multi.New(t))
	} else {
		if isShowConsole {
			log.SetHandler(multi.New(grayLog, t))
		} else {
			log.SetHandler(multi.New(grayLog))
		}
		level := viper.GetInt("logger.handlers.graylog.level")
		log.SetLevel(log.Level(level))
	}
	fields := make(log.Fields)
	grayFields := viper.GetString("logger.fields")
	grayFieldsArray := strings.Split(grayFields, ",")

	if len(grayFieldsArray) > 0 {
		for i := 0; i < len(grayFieldsArray); i++ {
			temp := strings.Split(grayFieldsArray[i], ":")
			if len(temp) > 1 {
				fields[string(temp[0])] = temp[1]
			}
		}
	}
	if newFields != nil {
		for k, v := range newFields[0] {
			fields[k] = v
		}
	}
	return log.WithFields(fields)

}

func getGrayLog() log.Handler {
	upd := viper.GetString("logger.handlers.graylog.protocol") + "://" +
		viper.GetString("logger.handlers.graylog.host") + ":" +
		viper.GetString("logger.handlers.graylog.port")
	fmt.Println(upd)
	e, err := graylog.New(upd)
	if err != nil {
		return &l
	}
	return e
}

func GinLogger(c *gin.Context) *log.Entry {
	if l, ok := c.Get("logger"); ok {
		return l.(*log.Entry)
	} else {
		return GrayLog()
	}
}
