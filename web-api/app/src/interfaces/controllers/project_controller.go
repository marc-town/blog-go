package controllers

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/marc-town/blog-go/web-api/app/domain/model"
	"github.com/marc-town/blog-go/web-api/app/interfaces/database"
	"github.com/marc-town/blog-go/web-api/app/usecase"
)

type ProjectController struct {
	Interactor usecase.ProjectInteractor
}

func NewProjectController(sqlHandler database.SqlHandler) *ProjectController {
	return &ProjectController{
		Interactor: usecase.ProjectInteractor{
			ProjectRepository: &database.ProjectRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ProjectController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	project, err := controller.Interactor.ProjectById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, SetResponse(project))
	return
}

func (controller *ProjectController) Index(c echo.Context) (err error) {
	projects, err := controller.Interactor.Projects()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, SetResponse(projects))
	return
}

func (controller *ProjectController) Create(c echo.Context) (err error) {
	a := model.Project{}
	c.Bind(&a)
	project, err := controller.Interactor.Add(a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, SetResponse(project))
	return
}

func (controller *ProjectController) Save(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	a := model.Project{}
	c.Bind(&a)
	a.ID = id
	project, err := controller.Interactor.Update(a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, SetResponse(project))
	return
}

func (controller *ProjectController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	project := model.Project{
		ID: id,
	}
	err = controller.Interactor.DeleteById(project)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, SetResponse(project))
	return
}
