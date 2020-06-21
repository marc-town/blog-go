package infrastructure

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/marc-town/blog-go/web-api/app/interfaces/controllers"
)

func Init() {
	log.Print("Init router")
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Sample
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is Running!")
	})

	// Global prefix
	g := e.Group("api/v1")
	// Controllers
	articlesController := controllers.NewArticleController(NewSqlHandler())
	// historyController := controllers.HistoryController(NewSqlHandler())

	// Routting endpoint
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

	// Start erver
	e.Logger.Fatal(e.Start(":1323"))
}
