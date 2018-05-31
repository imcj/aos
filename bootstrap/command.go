package bootstrap

import (
	"github.com/spf13/viper"
	"fmt"
	// "github.com/aos-stack/env"
)



// type EnvCommand struct {}

// func (c EnvCommand)Execute() {
// 	struct 
// 	env.SetProvider(provider)
// }

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
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// viper.SetEnvPrefix(cmdRoot)
	// replacer := strings.NewReplacer(".", "_")
	// viper.SetEnvKeyReplacer(replacer)
}