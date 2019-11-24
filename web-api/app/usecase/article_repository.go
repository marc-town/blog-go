package usecase

import "github.com/marc-town/blog-go/web-api/api/domain"

type ArticleRepository interface {
	FindById(id int) (domain.Article, error)
	FindAll() (domain.Articles, error)
	Store(domain.Article) (domain.Article, error)
	Update(domain.Article) (domain.Article, error)
	DeleteById(domain.Article) error
}