package controllers

import (
  "strconv"

  "github.com/marc-town/blog-go/web-api/api/domain"
  "github.com/marc-town/blog-go/web-api/api/database"
  "github.com/marc-town/blog-go/web-api/api/usecase"
)

type ArticleController struct {
  Interactor usecase.ArticleInteractor
}

func NewArticleController(sqlHandler database.SqlHandler) *ArticleController {
  return &ArticleController{
    Interactor: usecase.UserInteractor{
      ArticleRepository: &database.ArticleRepository{
        SqlHandler: sqlHandler,
      }
    }
  }
}

func (controller *ArticleController) Show(c echo.Context) (err error) {
  id, _ := strconv.Atoi(c.Param("id"))
  article, err := controller.Interactor.ArticleById(id)
  if err != nil {
    c.JSON(500, NewError(err)
    return
  }
  c.JSON(200, article)
  return
}

func (controller *ArticleController) Index(c echo.Context) (err error) {
  articles, err := controller.Interactor.Articles()
  if err != nil {
    c.JSON(500, NewError(err))
    return
  }
  c.JSON(200, articles)
  return
}

func (controller *ArticleController) Create(c echo.Context) (err error) {
  a := domain.Article{}
  c.Bind(&a)
  article, err := controller.Interactor.Add(a)
  if err != nil {
    c.JSON(500, NewError(err))
    return
  }
  c.JSON(201, article)
  return
}

func (controller *ArticleController) Save(c echo.Context) (err error) {
  a := domain.Article{}
  c.Bind(&a)
  article, err := controller.Interactor.Update(a)
  if err != nil {
    c.JSON(500, NewError(err))
    return
  }
  c.JSON(201, article)
  return
}

func (controller *ArticleController) Delete(c echo.Context) (err error) {
  id, _ := strconv.Atoi(c.Param("id"))
  article := domain.Article{
    ID: id,
  }
  err = controller.Interactor.DeleteById(article)
  if err != nil {
    c.JSON(500, NewError(err))
    return
  }
  c.JSON(200, article)
  return
}