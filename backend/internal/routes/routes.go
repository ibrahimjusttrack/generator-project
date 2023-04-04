package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myapp/internal/handlers"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", handlers.Hello)
	e.GET("/templates/all", handlers.GetAllTemplates)
	e.GET("/fields/:id", handlers.GetFields)
	e.POST("/template", handlers.CreateTemplate)
	e.POST("/upload-template/:id", handlers.UploadTemplate)
	e.POST("/field/:id", handlers.CreateField)
	e.POST("/config/:id", handlers.CreateJSONConfig)
	return e
}
