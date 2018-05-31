package bootstrap

import (
	// "fmt"
	"github.com/spf13/viper"
	"github.com/aos-stack/aos/interfaces"
)

type BootstrapRunCommand struct {
}

func (c BootstrapRunCommand)Execute() {
	cmd := BootstrapConfigCommand{ConfigPath: "conf/bootstrap.json"}
	cmd.Execute()
	
	viper.AddConfigPath("conf")
	interfaces.AddCommand("RemoteConfigCommand", RemoteConfigCommand{})
	interfaces.AddCommand("ConfigCommand", ConfigCommand{})
	interfaces.AddCommand("RedisCommand", RedisCommand{})
	interfaces.AddCommand("HTTPServerCommand", HTTPServerCommand{})
	interfaces.ExecuteCommands(Config.Application.Cycle)
}