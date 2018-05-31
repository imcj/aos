package utils

import (
	"github.com/aos-stack/aos/pkg/dbconf"
	"github.com/aos-stack/aos/pkg/setting"
	"time"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"io"
)

const (
	// DB pool config
	MaxIdleConns    int           = 50
	MaxOpenConns    int           = 200
	ConnMaxLifetime time.Duration = 1 * time.Hour
)

// 高顿库
var DbOne = NewEngine(dbconf.GDDBConf[0].DriverName, dbconf.GDDBConf[0].DriverDns)

type Db struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	driverName      string
	dataSourceName  string
	out             io.Writer // 写日志
	e               *xorm.Engine
}

func NewEngine(driverName string, dataSourceName string) *xorm.Engine {
	var d = Db{MaxIdleConns: MaxIdleConns,
		MaxOpenConns:    MaxOpenConns,
		ConnMaxLifetime: ConnMaxLifetime,
		dataSourceName:  dataSourceName,
		driverName:      driverName,
		out:             GraySql,
	}
	RetryLog("start : " + dataSourceName)

	d.InitEngine()
	d.setPoolNum()
	d.Ping()

	return d.e
}

func (d *Db) setPoolNum() {
	dbLogger := xorm.NewSimpleLogger(d.out)
	dbLogger.ShowSQL(true)
	dbLogger.SetLevel(core.LOG_INFO)
	d.e.Logger().SetLevel(core.LOG_INFO)
	d.e.SetLogger(dbLogger)
	d.e.ShowSQL(true)
	d.e.ShowExecTime(true)
	d.e.DB().SetConnMaxLifetime(d.ConnMaxLifetime)
	d.e.SetMaxIdleConns(d.MaxIdleConns)
	d.e.SetMaxOpenConns(d.MaxOpenConns)
}

func (gd *Db) InitEngine() error {
	e, err := xorm.NewEngine(gd.driverName, gd.dataSourceName)
	if err != nil {
		RetryLog("db_err : " + err.Error() + gd.dataSourceName)
		fmt.Println("db_err : " + err.Error() + gd.dataSourceName)
		panic(err)
	}
	gd.e = e
	return nil
}

// 定时 ping 数据库状态
func (d *Db) Ping() {
	go func() {
		var i time.Duration = 0
		for {
			if err := d.e.Ping(); err != nil {
				i++
				RetryLog("db_err_ping() err : %s, num: %d ", err.Error(), i)
				d.InitEngine()
				d.setPoolNum()
				time.Sleep(i * 200 * time.Millisecond) // 200 毫秒
			} else {
				i = 0
				time.Sleep(5 * time.Minute)
			}
		}
	}()
}

func RetryLog(format string, v ...interface{}) {
	setting.Logger.Infof(format, v)
}
