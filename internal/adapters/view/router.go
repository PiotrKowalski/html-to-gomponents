package view

import (
	"github.com/labstack/echo/v4"
	"html-to-gomponents/internal/app"
)

func NewRouter(e *echo.Group, app app.Application) *echo.Group {

	e.GET("", createIndexPageHandler())
	e.POST("/parse", createParseHandler(app))
	return e
}
