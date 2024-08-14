package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html-to-gomponents/internal/adapters/view"
	"html-to-gomponents/internal/app"
)

func New(app app.Application) (*echo.Echo, error) {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "uri=${uri}, " +
			"status=${status}, " +
			"request_id=${id}, " +
			"remote_ip=${remote_ip}" +
			"latency=${latency_human}\n",
	}))
	e.Use(middleware.RequestID())

	view.NewRouter(e.Group(""), app)
	return e, nil
}
