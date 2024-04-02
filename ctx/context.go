package ctx

import (
	"context"

	"github.com/asdgo/asdgo/database"

	"github.com/labstack/echo/v4"
)

type ContextKey string

var CsrfTokenKey = ContextKey("csrfToken")
var UserIDKey = ContextKey("userID")

func CsrfToken(ctx echo.Context) string {
	if csrfToken, ok := ctx.Get(string(CsrfTokenKey)).(string); ok {
		return csrfToken
	}

	return ""
}

func UserIsAuthenticated[T any](ctx T) bool {
	return UserID[T](ctx) != ""
}

func UserID[T any](ctx T) string {
	if con, ok := any(ctx).(echo.Context); ok {
		if userID, ok := con.Get(string(UserIDKey)).(string); ok {
			return userID
		}
	}

	if con, ok := any(ctx).(context.Context); ok {
		if userID, ok := con.Value(UserIDKey).(string); ok {
			return userID
		}
	}

	return ""
}

func User(ctx echo.Context) database.User {
	var user database.User
	database.Instance.Find(&user, "id = ?", UserID(ctx))

	return user
}
