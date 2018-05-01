package persistence

import (
	"aos/project/domain"

	"github.com/go-xorm/xorm"
)

type ProjectDAODBAL struct {
	Table domain.ProjectModel
	domain.ProjectDAO
	Engine *xorm.Engine
}

func (c *ProjectDAODBAL) List(b int) domain.ProjectModel {

	var bp domain.ProjectModel
	c.Engine.Table("gd_project").Where("id = 8").Get(&bp)

	return bp
}
