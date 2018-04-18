package setting

import (
	"os"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/graylog"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

var Logger *log.Entry

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		// log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {

}

func GrayLog(newFields ...map[string]interface{}) *log.Entry {
	graylogInfo, _ := Cfg.GetSection("log")
	e, _ := graylog.New(graylogInfo.Key("LOG_UDP").MustString("udp://g02.graylog.gaodunwangxiao.com:5504"))
	isShowConsole := graylogInfo.Key("IS_SHOW_CONSOLE").MustBool(false)
	if isShowConsole {
		t := text.New(os.Stderr)
		log.SetHandler(multi.New(e, t))
	} else {
		log.SetHandler(multi.New(e))
	}

	fields := make(log.Fields)
	grayFields := graylogInfo.Key("LOG_FIELDS").MustString("item:ginlog")
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
	level := graylogInfo.Key("LOG_LEVEL").MustInt(-1)
	log.SetLevel(log.Level(level))
	return log.WithFields(fields)
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		// Logger.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		// Logger.Fatalf("Fail to get section 'app': %v", err)
	}

	// JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
