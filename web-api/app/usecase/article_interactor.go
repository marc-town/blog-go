package usecase

import (
	"github.com/marc-town/blog-go/web-api/app/domain/model"
	"github.com/marc-town/blog-go/web-api/app/domain/repository"
)

type ArticleInteractor struct {
	ArticleRepository repository.ArticleRepository
}

func (interactor *ArticleInteractor) ArticleById(id int) (article model.Article, err error) {
	article, err = interactor.ArticleRepository.FindById(id)
	return
}

func (interactor *ArticleInteractor) Articles() (articles model.Articles, err error) {
	articles, err = interactor.ArticleRepository.FindAll()
	return
}

func (interactor *ArticleInteractor) Add(a model.Article) (article model.Article, err error) {
	article, err = interactor.ArticleRepository.Store(a)
	return
}

func (interactor *ArticleInteractor) Update(a model.Article) (article model.Article, err error) {
	article, err = interactor.ArticleRepository.Update(a)
	return
}

func (interactor *ArticleInteractor) DeleteById(a model.Article) (err error) {
	err = interactor.ArticleRepository.DeleteById(a)
	return
}
