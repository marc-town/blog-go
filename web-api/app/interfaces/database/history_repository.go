package database

import "github.com/marc-town/blog-go/web-api/app/domain/model"

type HistoryRepository struct {
	SqlHandler
}

func (repo *HistoryRepository) FindById(id int) (history model.History, err error) {
	if err = repo.Find(&history, id).Error; err != nil {
		return
	}
	return
}

func (repo *HistoryRepository) FindAll() (histories model.Histories, err error) {
	if err = repo.Find(&histories).Error; err != nil {
		return
	}
	return
}

func (repo *HistoryRepository) Store(a model.History) (history model.History, err error) {
	if err = repo.Create(&a).Error; err != nil {
		return
	}
	history = a
	return
}

func (repo *HistoryRepository) Update(a model.History) (history model.History, err error) {
	if err = repo.Save(&a).Error; err != nil {
		return
	}
	history = a
	return
}

func (repo *HistoryRepository) DeleteById(history model.History) (err error) {
	if err = repo.Delete(&history).Error; err != nil {
		return
	}
	return
}
