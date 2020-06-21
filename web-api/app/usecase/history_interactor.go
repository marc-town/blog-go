package usecase

import (
	"github.com/marc-town/blog-go/web-api/app/domain/model"
	"github.com/marc-town/blog-go/web-api/app/domain/repository"
)

type HistoryInteractor struct {
	HistoryRepository repository.HistoryRepository
}

func (interactor *HistoryInteractor) HistoryById(id int) (history model.History, err error) {
	history, err = interactor.HistoryRepository.FindById(id)
	return
}

func (interactor *HistoryInteractor) Histories() (histories model.Histories, err error) {
	histories, err = interactor.HistoryRepository.FindAll()
	return
}

func (interactor *HistoryInteractor) Add(a model.History) (history model.History, err error) {
	history, err = interactor.HistoryRepository.Store(a)
	return
}

func (interactor *HistoryInteractor) Update(a model.History) (history model.History, err error) {
	history, err = interactor.HistoryRepository.Update(a)
	return
}

func (interactor *HistoryInteractor) DeleteById(a model.History) (err error) {
	err = interactor.HistoryRepository.DeleteById(a)
	return
}
