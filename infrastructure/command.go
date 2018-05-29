package infrastructure

import (
	"github.com/spf13/viper"
	"fmt"
)

type RemoteConfigCommand struct {}

func (r RemoteConfigCommand)Execute() {

}

type ConfigCommand struct {

}

func (c ConfigCommand)Execute() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("../conf") // 目的是让测试通过
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}