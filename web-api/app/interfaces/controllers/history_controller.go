package controllers

import (
	"log"
	"strconv"

	"github.com/labstack/echo"
	"github.com/marc-town/blog-go/web-api/app/domain/model"
	"github.com/marc-town/blog-go/web-api/app/interfaces/database"
	"github.com/marc-town/blog-go/web-api/app/usecase"
)

type HistoryController struct {
	Interactor usecase.HistoryInteractor
}

func NewHistoryController(sqlHandler database.SqlHandler) *HistoryController {
	return &HistoryController{
		Interactor: usecase.HistoryInteractor{
			HistoryRepository: &database.HistoryRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *HistoryController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	history, err := controller.Interactor.HistoryById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, SetResponse(history))
	return
}

func (controller *HistoryController) Index(c echo.Context) (err error) {
	log.Print("Init history controller")
	histories, err := controller.Interactor.Histories()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, SetResponse(histories))
	return
}

func (controller *HistoryController) Create(c echo.Context) (err error) {
	a := model.History{}
	c.Bind(&a)
	history, err := controller.Interactor.Add(a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, SetResponse(history))
	return
}

func (controller *HistoryController) Save(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	a := model.History{}
	c.Bind(&a)
	a.ID = id
	history, err := controller.Interactor.Update(a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, SetResponse(history))
	return
}

func (controller *HistoryController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	history := model.History{
		ID: id,
	}
	err = controller.Interactor.DeleteById(history)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, SetResponse(history))
	return
}
