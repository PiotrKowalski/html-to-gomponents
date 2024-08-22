package router

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const expiryTIme = 15

func sessionIdentifier(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_id")
		if err != nil || cookie.Value == "" {
			sessionID := uuid.New().String()

			cookie = &http.Cookie{
				Name:    "session_id",
				Value:   sessionID,
				Path:    "/",
				Expires: time.Now().Add(time.Duration(expiryTIme) * time.Minute),
			}
			c.SetCookie(cookie)
		} else {
			cookie.Expires = time.Now().Add(time.Duration(expiryTIme) * time.Minute)
			c.SetCookie(cookie)
		}
		c.Set("session_id", cookie.Value)

		return next(c)
	}
}
