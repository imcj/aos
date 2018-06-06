package bootstrap

import (
	"fmt"
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

		containerDatabases[name] = engine
	}
	aos.ContainerSet("database", containerDatabases)
	
	keepAlive()
}

func keepAlive() {
	database := aos.ContainerGet("database").(map[string]*xorm.Engine)
	engine := database["default"]
	err := engine.Ping()
}