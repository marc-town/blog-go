package database

import "github.com/marc-town/blog-go/web-api/app/domain/model"

type ProjectRepository struct {
	SqlHandler
}

func (repo *ProjectRepository) FindById(id int) (project model.Project, err error) {
	if err = repo.Find(&project, id).Error; err != nil {
		return
	}
	return
}

func (repo *ProjectRepository) FindAll() (projects model.Projects, err error) {
	if err = repo.Find(&projects).Error; err != nil {
		return
	}
	return
}

func (repo *ProjectRepository) Store(a model.Project) (project model.Project, err error) {
	if err = repo.Create(&a).Error; err != nil {
		return
	}
	project = a
	return
}

func (repo *ProjectRepository) Update(a model.Project) (project model.Project, err error) {
	if err = repo.Save(&a).Error; err != nil {
		return
	}
	project = a
	return
}

func (repo *ProjectRepository) DeleteById(project model.Project) (err error) {
	if err = repo.Delete(&project).Error; err != nil {
		return
	}
	return
}
