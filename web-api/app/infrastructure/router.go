package infrastructure

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/marc-town/blog-go/web-api/app/interfaces/controllers"
)

func Init() {
	// Echo instance
	e := echo.New()

	// Controllers
	articlesController := controllers.NewArticleController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Sample
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, GOlang!")
	})

	// Routting endpoint
	e.GET("/articles", func(c echo.Context) error { return articlesController.Index(c) })
	e.GET("/article/:id", func(c echo.Context) error { return articlesController.Show(c) })
	e.POST("/article", func(c echo.Context) error { return articlesController.Create(c) })
	e.PUT("/article/:id", func(c echo.Context) error { return articlesController.Save(c) })
	e.DELETE("/article/:id", func(c echo.Context) error { return articlesController.Delete(c) })

	// Start erver
	e.Logger.Fatal(e.Start(":1323"))
}
