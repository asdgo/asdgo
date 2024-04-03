package asession

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type Session struct {
	Name string

	store *sessions.CookieStore
}

var Instance *Session

func New(name string) {
	store := sessions.NewCookieStore([]byte(os.Getenv("APP_KEY")))
	store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   86400 * 7,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}

	if name == "" {
		name = "asdgo_session"
	}

	Instance = &Session{
		Name:  name,
		store: store,
	}
}

func (s *Session) Store() *sessions.CookieStore {
	return s.store
}

func (s *Session) Get(c echo.Context, key string) string {
	sess, _ := s.store.Get(c.Request(), s.Name)

	if _, ok := sess.Values[key]; !ok {
		return ""
	}

	if ok := sess.Values[key].(string); ok != "" {
		return ok
	}

	return ""
}

func (s *Session) Set(c echo.Context, key string, value string) {
	sess, _ := s.store.Get(c.Request(), s.Name)
	sess.Values[key] = value

	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		fmt.Printf("[SESSION ERROR]: %v\n", err)
	}
}

func (s *Session) Delete(c echo.Context, key string) {
	sess, _ := s.store.Get(c.Request(), s.Name)
	sess.Options.MaxAge = -1

	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		fmt.Printf("[SESSION ERROR]: %v\n", err)
	}
}
