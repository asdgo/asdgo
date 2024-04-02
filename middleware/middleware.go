package middleware

import (
	"net/http"

	"github.com/asdgo/asdgo/ctx"

	"github.com/labstack/echo/v4"
)

func Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !ctx.UserIsAuthenticated(c) {
				return c.Redirect(http.StatusFound, "/auth/login")
			}

			return next(c)
		}
	}
}
