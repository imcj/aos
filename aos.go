package aos

import (
	"github.com/aos-stack/aos/bootstrap"
)

func Run() {
	bootstrap := bootstrap.BootstrapRunCommand{}
	bootstrap.Execute()
}