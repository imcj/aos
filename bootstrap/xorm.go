package bootstrap

import (
	"fmt"
	"time"
	"github.com/spf13/viper"
	// "github.com/aos-stack/aos/interfaces"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	aos "github.com/aos-stack/aos/container"
	"github.com/aos-stack/aos/config"
)

type XORMCommand struct {
}

func (c XORMCommand)Execute() {
	databases := viper.Get("database").(map[string]interface{})
	assembler := &bootstrap.ConfigAssembler{}

	containerDatabases := make(map[string]*xorm.Engine)

	for name, database := range databases {
		config := assembler.Database(database.(map[string]interface{}))
		source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.Username, config.Password, config.Host, config.Port, config.Database)
		engine, err := xorm.NewEngine(config.Type, source)
		if err != nil {
			// RetryLog("db_err : " + err.Error() + source)
			fmt.Println("db_err : " + err.Error() + source)
			panic(err)
		}

		// TODO:
		
		// dbLogger := xorm.NewSimpleLogger()
		// dbLogger.ShowSQL(config.LogSQL)
		// dbLogger.SetLevel(core.LogLevel)
		// engine.Logger().SetLevel(config.LogLevel)
		// engine.SetLogger(dbLogger)
		engine.ShowSQL(config.LogSQL)
		engine.ShowExecTime(config.LogSQLExecuteTime)
		engine.DB().SetConnMaxLifetime(config.MaxLeftTime)
		engine.SetMaxIdleConns(config.MaxIdle)
		engine.SetMaxOpenConns(config.MaxOpen)

		containerDatabases[name] = engine
	}
	aos.ContainerSet("database", containerDatabases)
	
	go keepAlive()
}

func keepAlive() {
	database := aos.ContainerGet("database").(map[string]*xorm.Engine)

	for {
		for _, engine := range database {
			err := engine.Ping()

			if err != nil {
				fmt.Println("Ping error")
			}
		}
		time.Sleep(1 * time.Minute)
	}
}