package router

import (
	"context"
	"net/http"

	"github.com/asdgo/asdgo/session"
	"github.com/asdgo/asdgo/template"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
)

type Router struct {
	*chi.Mux
}

var Instance *Router

func New() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Use(func(next http.Handler) http.Handler {
		csrfHandler := nosurf.New(next)
		csrfHandler.SetBaseCookie(http.Cookie{
			Name:     "asdgo_csrf_token",
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
			ctx := context.WithValue(r.Context(), template.CsrfTokenKey, nosurf.Token(r))
			ctx = context.WithValue(ctx, template.UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	Instance = &Router{
		r,
	}
}
