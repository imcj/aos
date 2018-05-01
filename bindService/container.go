package container

import (
	"aos/controller"
	"aos/pkg/utils"
	"aos/project/domain"
	"aos/project/infrastructure/persistence/dbal"
	"aos/project/service"

	"github.com/go-xorm/xorm"
)

//初始化服务于绑定，顺序要不要注意》？
var dbConnect xorm.Engine = initEng(0)
var projectDAO domain.ProjectDAO = NewProjectDAODBAL(dbConnect)
var ProjectService service.ProjectService = service.NewProjectServiceImpl(projectDAO)
var ProjectController *controller.TestApi = controller.NewProjectController(ProjectService)

func initEng(num int) xorm.Engine {
	eng, _ := utils.InitEng(num)
	return *eng
}

//方法可以写在这里也可以放在具体的内部 采用哪个？
func NewProjectDAODBAL(eng xorm.Engine) *persistence.ProjectDAODBAL {
	var a *persistence.ProjectDAODBAL = &persistence.ProjectDAODBAL{}
	a.Engine = &eng
	return a
}
