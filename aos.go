package aos

import (
	"github.com/aos-stack/aos/bootstrap"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(registry func(engine *gin.Engine)) {
	bootstrap.Registry = registry
}

func Run() {
	bootstrap := bootstrap.BootstrapRunCommand{}
	bootstrap.Execute()
}