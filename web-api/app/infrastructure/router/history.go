package router

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marc-town/blog-go/web-api/app/infrastructure"
	"github.com/marc-town/blog-go/web-api/app/interfaces/controllers"
)

func HistoryRouter(g *echo.Group) {
	log.Print("History roter")

	// Controllers
	articlesController := controllers.NewArticleController(infrastructure.NewSqlHandler())
	// historyController := controllers.HistoryController(NewSqlHandler())

	g.GET("/histories", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "histories")
	})
	// routting endpoint
	g.GET("/articles", func(c echo.Context) error { return articlesController.Index(c) })
	g.GET("/article/:id", func(c echo.Context) error { return articlesController.Show(c) })
	g.POST("/article", func(c echo.Context) error { return articlesController.Create(c) })
	g.PUT("/article/:id", func(c echo.Context) error { return articlesController.Save(c) })
	g.DELETE("/article/:id", func(c echo.Context) error { return articlesController.Delete(c) })

	// e.GET("/histories", func(c echo.Context) error { return historyController.Index(c) })
	// e.GET("/histories/:id", func(c echo.Context) error { return historyController.Show(c) })
	// e.POST("/histories", func(c echo.Context) error { return historyController.Create(c) })
	// e.PUT("/histories/:id", func(c echo.Context) error { return historyController.Save(c) })
	// e.DELETE("/histories/:id", func(c echo.Context) error { return historyController.Delete(c) })
}
