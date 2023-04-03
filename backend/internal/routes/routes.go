package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/handlers"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", handlers.Hello)
	e.GET("/templates/all", handlers.GetAllTemplates)

	return e
}
