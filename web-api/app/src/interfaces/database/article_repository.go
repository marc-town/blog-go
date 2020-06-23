package database

import "github.com/marc-town/blog-go/web-api/app/domain/model"

type ArticleRepository struct {
	SqlHandler
}

func (repo *ArticleRepository) FindById(id int) (article model.Article, err error) {
	if err = repo.Find(&article, id).Error; err != nil {
		return
	}
	return
}

func (repo *ArticleRepository) FindAll() (articles model.Articles, err error) {
	if err = repo.Find(&articles).Error; err != nil {
		return
	}
	return
}

func (repo *ArticleRepository) Store(a model.Article) (article model.Article, err error) {
	if err = repo.Create(&a).Error; err != nil {
		return
	}
	article = a
	return
}

func (repo *ArticleRepository) Update(a model.Article) (article model.Article, err error) {
	if err = repo.Save(&a).Error; err != nil {
		return
	}
	article = a
	return
}

func (repo *ArticleRepository) DeleteById(article model.Article) (err error) {
	if err = repo.Delete(&article).Error; err != nil {
		return
	}
	return
}
