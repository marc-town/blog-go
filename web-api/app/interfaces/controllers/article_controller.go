package controllers

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/marc-town/blog-go/web-api/app/domain/model"
	"github.com/marc-town/blog-go/web-api/app/interfaces/database"
	"github.com/marc-town/blog-go/web-api/app/usecase"
)

type ArticleController struct {
	Interactor usecase.ArticleInteractor
}

func NewArticleController(sqlHandler database.SqlHandler) *ArticleController {
	return &ArticleController{
		Interactor: usecase.ArticleInteractor{
			ArticleRepository: &database.ArticleRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ArticleController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := controller.Interactor.ArticleById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, SetResponse(article))
	return
}

func (controller *ArticleController) Index(c echo.Context) (err error) {
	articles, err := controller.Interactor.Articles()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, SetResponse(articles))
	return
}

func (controller *ArticleController) Create(c echo.Context) (err error) {
	a := model.Article{}
	c.Bind(&a)
	article, err := controller.Interactor.Add(a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, SetResponse(article))
	return
}

func (controller *ArticleController) Save(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	a := model.Article{}
	c.Bind(&a)
	a.ID = id
	article, err := controller.Interactor.Update(a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, SetResponse(article))
	return
}

func (controller *ArticleController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	article := model.Article{
		ID: id,
	}
	err = controller.Interactor.DeleteById(article)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, SetResponse(article))
	return
}
