package infrastructure

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/aos-stack/aos/interfaces"
	"github.com/aos-stack/aos/infrastructure"
)

func TestConfigCommand(t * testing.T) {
	interfaces.AddCommand(infrastructure.RemoteConfigCommand{})
	interfaces.AddCommand(infrastructure.ConfigCommand{})
	
	interfaces.ExecuteCommands()
	assert.Equal(t, 1, 1)
}