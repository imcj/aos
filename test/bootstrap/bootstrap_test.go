package bootstrap

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/aos-stack/aos/bootstrap"
)

func TestBootstrapConfigCommand(t * testing.T) {
	cmd := bootstrap.BootstrapConfigCommand{ConfigPath: "../../conf/bootstrap.json"}
	cmd.Execute()
	assert.Equal(t, "", bootstrap.Config.Consul)
}