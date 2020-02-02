package repository

import "github.com/marc-town/blog-go/web-api/app/domain/model"

type ArticleRepository interface {
	FindById(id int) (model.Article, error)
	FindAll() (model.Articles, error)
	Store(model.Article) (model.Article, error)
	Update(model.Article) (model.Article, error)
	DeleteById(model.Article) error
}
