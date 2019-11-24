package usecase

import "github.com/marc-town/blog-go/web-api/api/domain"

type ArticleInteractor struct {
	ArticleRepository ArticleRepository
}

func (interactor *ArticleInteractor) ArticleById(id int) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.FindById(id)
	return
}

func (interactor *ArticleInteractor) Article() (articles domain.Articles, err error) {
	articles, err = interactor.ArticleRepository.FindAll()
	return
}

func (interactor *ArticleInteractor) Add(a domain.Article) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.Store(a)
	return
}

func (interactor *ArticleInteractor) Update(a domain.Article) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.Update(a)
	return
}

func (interactor *ArticleInteractor) DeleteById(a domain.Article) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.FindById(a)
	return
}