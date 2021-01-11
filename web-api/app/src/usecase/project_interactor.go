package usecase

import (
	"github.com/marc-town/blog-go/web-api/app/domain/model"
	"github.com/marc-town/blog-go/web-api/app/domain/repository"
)

type ProjectInteractor struct {
	ProjectRepository repository.ProjectRepository
}

func (interactor *ProjectInteractor) ProjectById(id int) (project model.Project, err error) {
	project, err = interactor.ProjectRepository.FindById(id)
	return
}

func (interactor *ProjectInteractor) Projects() (projects model.Projects, err error) {
	projects, err = interactor.ProjectRepository.FindAll()
	return
}

func (interactor *ProjectInteractor) Add(a model.Project) (project model.Project, err error) {
	project, err = interactor.ProjectRepository.Store(a)
	return
}

func (interactor *ProjectInteractor) Update(a model.Project) (project model.Project, err error) {
	project, err = interactor.ProjectRepository.Update(a)
	return
}

func (interactor *ProjectInteractor) DeleteById(a model.Project) (err error) {
	err = interactor.ProjectRepository.DeleteById(a)
	return
}
