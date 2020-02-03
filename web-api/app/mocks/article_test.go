package mocks

import (
	reflect "reflect"
	"testing"
	model "github.com/marc-town/blog-go/web-api/app/domain/model"
	"github.com/marc-town/blog-go/web-api/app/usecase"

	"github.com/golang/mock/gomock"
)

func TestView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected []*model.Article
	var err error

	mockSample := NewMockTodoRepository(ctrl)
	mockSample.EXPECT().FindAll().Return(expected, err)

	// mockを利用してarticle_interactor.Articles()をテストする
	// u := usecase.ArticleInteractor()
	// result, err := usecase.
}