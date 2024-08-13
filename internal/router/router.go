package router

import (
	"github.com/labstack/echo/v4"
	"html-to-gomponents/internal/adapters/view"
	"html-to-gomponents/internal/app"
)

func New(app app.Application) (*echo.Echo, error) {
	e := echo.New()

	view.NewRouter(e.Group(""), app)
	return e, nil
}
