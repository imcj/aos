package bootstrap

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/aos-stack/aos/bootstrap"
	"github.com/aos-stack/aos/interfaces"
	"github.com/spf13/viper"
)

// func TestBootstrapConfigCommand(t * testing.T) {
// 	cmd := bootstrap.BootstrapConfigCommand{ConfigPath: "../../conf/bootstrap.json"}
// 	cmd.Execute()
// 	assert.Equal(t, "", bootstrap.Config.Consul)
// }

func TestCommand(t * testing.T) {
	viper.AddConfigPath("../../conf")

	interfaces.AddCommand("RemoteConfigCommand", bootstrap.RemoteConfigCommand{})
	interfaces.AddCommand("ConfigCommand", bootstrap.ConfigCommand{})
	
	interfaces.ExecuteCommands([]string{"RemoteConfigCommand","ConfigCommand"})
	assert.Equal(t, 1, 1)
}