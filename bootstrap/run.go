package bootstrap

import (
	// "fmt"

	"github.com/aos-stack/aos/interfaces"
	"github.com/spf13/viper"
)

type BootstrapRunCommand struct {
}

func (c BootstrapRunCommand) Execute() {
	cmd := BootstrapConfigCommand{ConfigPath: "conf/bootstrap.json"}
	cmd.Execute()

	viper.AddConfigPath("conf")

	interfaces.AddCommand("RemoteConfigCommand", RemoteConfigCommand{})
	interfaces.AddCommand("ConfigCommand", ConfigCommand{})
	interfaces.AddCommand("ConsulCommand", ConsulCommand{})
	interfaces.AddCommand("LoggerCommand", LoggerCommand{})
	interfaces.AddCommand("XORMCommand", XORMCommand{})
	interfaces.AddCommand("RedisCommand", RedisCommand{})
	interfaces.AddCommand("GinHTTPMiddlewareCommand", GinHTTPMiddlewareCommand{})
	interfaces.AddCommand("HTTPServerCommand", HTTPServerCommand{})
	interfaces.ExecuteCommands(Config.Application.Cycle)
}
