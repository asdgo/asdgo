package acontext

import (
	"context"

	"github.com/asdgo/asdgo/adatabase"

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
	return UserID(ctx) != ""
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

func User[T any](ctx T) adatabase.User {
	var user adatabase.User
	adatabase.Instance.Find(&user, "id = ?", UserID(ctx))

	return user
}
