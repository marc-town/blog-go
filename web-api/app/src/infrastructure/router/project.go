package router

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marc-town/blog-go/web-api/app/infrastructure"
	"github.com/marc-town/blog-go/web-api/app/interfaces/controllers"
)

func HistoryRouter(g *echo.Group) {
	log.Print("Project roter")

	// Controllers
	projectsController := controllers.NewProjectController(infrastructure.NewSqlHandler())
	// historyController := controllers.HistoryController(NewSqlHandler())

	g.GET("/projects", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "projects")
	})
	// routting endpoint
	// g.GET("/projects", func(c echo.Context) error { return projectsController.Index(c) })
	// g.GET("/project/:id", func(c echo.Context) error { return projectsController.Show(c) })
	// g.POST("/project", func(c echo.Context) error { return projectsController.Create(c) })
	// g.PUT("/project/:id", func(c echo.Context) error { return projectsController.Save(c) })
	// g.DELETE("/project/:id", func(c echo.Context) error { return projectsController.Delete(c) })
}
