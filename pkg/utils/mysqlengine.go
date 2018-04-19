package utils

import (
	"aos/pkg/dbconf"
	"aos/pkg/errors"
	"aos/pkg/setting"
	"strconv"
	"strings"
	"time"

	"github.com/bernos/go-retry"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

const (
	GAODUN        int = 0
	DB_LIST_COUNT int = 1
)

var EngineList, err = NewEngine()

var engineList = make(map[int]*xorm.Engine)

var loggerHandle = setting.GrayLog()

func GetDBEng(engine *xorm.Engine, databaseNum int) func() (interface{}, error) {
	return func() (interface{}, error) {
		err := engine.Ping()
		if err != nil {
			configInfo, _ := dbconf.GetMySqlConfig()
			engine, _ = xorm.NewEngine(configInfo[databaseNum].DriverName, configInfo[databaseNum].DriverDns)
			return nil, err
		}
		return nil, nil
	}
}

func RetryLog(format string, v ...interface{}) {
	loggerHandle.Infof(format, v)
}

var GraySql GrayXormSql

// 实现 xrom 打印日志接口
type GrayXormSql struct {
}

func (GrayXormSql) Write(p []byte) (n int, err error) {
	if strings.Contains(strings.ToUpper(string(p)), "[SQL]") {
		s := string(p)
		t := strings.Split(s, "took:")
		tst := strings.Trim(t[1], " ")
		tst = strings.Replace(tst, "ms\n", "", -1)
		sqlTime, _ := strconv.ParseFloat(tst, 10)
		m := map[string]interface{}{
			"user_sql":           s,
			"user_sql_exec_time": sqlTime,
		}
		setting.Logger.Infof(s, m)
	}
	return 0, nil
}

// 初始化返回 engine
var InitEng = func(databaseNum int) (*xorm.Engine, error) {
	r := retry.Retry(GetDBEng(EngineList[databaseNum], databaseNum),
		retry.MaxRetries(5),
		retry.BaseDelay(time.Millisecond*200),
		retry.Log(RetryLog))
	_, err := r()
	if err != nil {
		loggerHandle.Info(err.(error).Error())
		return EngineList[databaseNum], err.(error)
	}
	return EngineList[databaseNum], nil
}

func NewEngine() (map[int]*xorm.Engine, error) {

	configInfo, err := dbconf.GetMySqlConfig()
	if err != nil {
		return nil, errors.New(0, err.Error())
	}
	for i := 0; i < DB_LIST_COUNT; i++ {
		engineList[i], err = xorm.NewEngine(configInfo[i].DriverName, configInfo[i].DriverDns)
		RetryLog("start : " + configInfo[i].DriverDns)
		if err != nil {
			RetryLog("db err : " + err.Error() + configInfo[i].DriverName + configInfo[i].DriverDns)
			return engineList, err
		}

		dblogger := xorm.NewSimpleLogger(GraySql)
		dblogger.ShowSQL(true)
		dblogger.SetLevel(core.LOG_INFO)
		engineList[i].Logger().SetLevel(core.LOG_INFO)
		engineList[i].SetLogger(dblogger)
		engineList[i].ShowSQL(true)
		engineList[i].ShowExecTime(true)
		err = engineList[i].Ping()
		if err != nil {
			RetryLog("db ping() err : " + err.Error() + configInfo[i].DriverName + configInfo[i].DriverDns)
			return engineList, err
		}
		tmpDb := engineList[i].DB()
		tmpDb.SetConnMaxLifetime(1 * time.Hour)
		engineList[i].SetMaxIdleConns(50)
		engineList[i].SetMaxOpenConns(200)
	}
	return engineList, nil
}
