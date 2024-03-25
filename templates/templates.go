package templates

import (
	"context"

	"github.com/asdgo/asdgo/database"
)

type ContextKey string

var CsrfTokenKey = ContextKey("csrfToken")
var UserIDKey = ContextKey("userID")

func CsrfToken(ctx context.Context) string {
	if csrfToken, ok := ctx.Value(CsrfTokenKey).(string); ok {
		return csrfToken
	}

	return ""
}

func UserIsAuthenticated(ctx context.Context) bool {
	return UserID(ctx) != ""
}

func UserID(ctx context.Context) string {
	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		return userID
	}

	return ""
}

func User(ctx context.Context) database.User {
	var user database.User
	database.Instance.Find(&user, "id = ?", UserID(ctx))

	return user
}
