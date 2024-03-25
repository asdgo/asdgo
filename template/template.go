package template

import (
	"context"

	"github.com/asdgo/asdgo/database"
)

type ContextKey string

var CsrfTokenKey = ContextKey("csrfToken")
var UserIDKey = ContextKey("userID")

type Template struct{}

var Instance *Template

func New() {
	Instance = &Template{}
}

func (t *Template) CsrfToken(ctx context.Context) string {
	if csrfToken, ok := ctx.Value(CsrfTokenKey).(string); ok {
		return csrfToken
	}

	return ""
}

func (t *Template) UserIsAuthenticated(ctx context.Context) bool {
	return t.UserID(ctx) != ""
}

func (t *Template) UserID(ctx context.Context) string {
	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		return userID
	}

	return ""
}

func (t *Template) User(ctx context.Context) database.User {
	var user database.User
	database.Instance.Find(&user, "id = ?", t.UserID(ctx))

	return user
}
