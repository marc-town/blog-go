package repository

import "github.com/marc-town/blog-go/web-api/app/domain/model"

type HistoryRepository interface {
	FindById(id int) (model.History, error)
	FindAll() (model.Histories, error)
	Store(model.History) (model.History, error)
	Update(model.History) (model.History, error)
	DeleteById(model.History) error
}
