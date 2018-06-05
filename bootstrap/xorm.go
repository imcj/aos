package bootstrap

import (
	// "fmt"
	"github.com/spf13/viper"
	"github.com/aos-stack/aos/interfaces"
)

type XORMCommand struct {
}

func (c XORMCommand)Execute() {
	
	engine, err := xorm.NewEngine(gd.driverName, gd.dataSourceName)
	if err != nil {
		RetryLog("db_err : " + err.Error() + gd.dataSourceName)
		fmt.Println("db_err : " + err.Error() + gd.dataSourceName)
		panic(err)
	}
}