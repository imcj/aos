package interfaces

import (
	"testing"
	"github.com/aos-stack/aos/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestAddClosureCommand(t *testing.T) {
	changed := false
	interfaces.AddClosureCommand("TestCommand", func () {
		changed = true
	})
	interfaces.ExecuteCommands([]string{"TestCommand"})
	assert.Equal(t, true, changed)

} 