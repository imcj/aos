package aos

import (
	"testing"
	"github.com/aos-stack/aos/container"
	"github.com/stretchr/testify/assert"
	_ "."
)

type Example struct {
	example string
}

func TestSet(t *testing.T) {
	aos.ContainerSet("example", Example{"example"})
	example := aos.ContainerGet("example")
	assert.Equal(t, "example", example.(Example).example)

	ok := aos.ContainerHas("id")
	assert.Equal(t, false, ok)
}