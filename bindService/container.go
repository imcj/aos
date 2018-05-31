package container

import (
	"github.com/aos-stack/aos/pkg/utils"
	// "github.com/aos-stack/aos/project/domain"
	// "github.com/aos-stack/aos/project/infrastructure/persistence/dbal"
	"github.com/aos-stack/aos/controller"

	"github.com/go-xorm/xorm"
)

var containerInstance *Container

func GetContainer() *Container {
	if containerInstance == nil {
		containerInstance = &Container{}
		containerInstance.init()
	}
	return containerInstance
}

type Container struct {
	TestApi *controller.TestApi
}

func (this *Container) init() {
	//var dbConnect *xorm.Engine = initEng(0)
	this.TestApi = controller.NewDemoController()

}

func initEng(num int) *xorm.Engine {
	eng := utils.DbOne
	return eng
}
