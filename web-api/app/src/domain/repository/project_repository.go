package repository

import "github.com/marc-town/blog-go/web-api/app/domain/model"

type ProjectRepository interface {
	FindById(id int) (model.Project, error)
	FindAll() (model.Projects, error)
	Store(model.Project) (model.Project, error)
	Update(model.Project) (model.Project, error)
	DeleteById(model.Project) error
}
