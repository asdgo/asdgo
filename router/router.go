package router

import (
	"context"
	"net/http"

	"github.com/asdgo/asdgo/session"
	"github.com/asdgo/asdgo/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
)

var Instance *chi.Mux

func New() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	r.Use(func(next http.Handler) http.Handler {
		csrfHandler := nosurf.New(next)
		csrfHandler.SetBaseCookie(http.Cookie{
			Name:     "checkeroni_csrf_token",
			HttpOnly: true,
			Path:     "/",
			MaxAge:   86400,
			SameSite: http.SameSiteLaxMode,
			Secure:   true,
		})

		return csrfHandler
	})

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := session.Get(r, "user_id")
			ctx := context.WithValue(r.Context(), templates.CsrfTokenKey, nosurf.Token(r))
			ctx = context.WithValue(ctx, templates.UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	Instance = r
}
