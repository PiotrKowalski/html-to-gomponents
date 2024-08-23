package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"html-to-gomponents/internal/adapters/view"
	"html-to-gomponents/internal/app"
	"net/http"
	"time"
)

func New(app app.Application) (*echo.Echo, error) {

	e := echo.New()
	e.Use(sessionIdentifier)

	e.Use(middleware.BodyLimit("25000"))

	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(10), Burst: 10, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	e.Use(middleware.RateLimiterWithConfig(config))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "uri=${uri}, " +
			"status=${status}, " +
			"request_id=${id}, " +
			"latency=${latency_human}, " +
			"session_id=${cookie:session_id}, " +
			"user_agent=${user_agent}, " +
			"bytes_in=${bytes_in}, " +
			"bytes_out=${bytes_out}\n",
	}))
	e.Use(middleware.RequestID())

	view.NewRouter(e.Group(""), app)
	return e, nil
}
