package routes

import (
	"github.com/labstack/echo"
	"myapp/api/handlers"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", handlers.Hello)

	return e
}
