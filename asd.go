package asdgo

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/asdgo/asdgo/ctx"
	"github.com/asdgo/asdgo/database"
	"github.com/asdgo/asdgo/hash"
	"github.com/asdgo/asdgo/mail"
	"github.com/asdgo/asdgo/session"
	"github.com/asdgo/asdgo/template"
	"github.com/asdgo/asdgo/validate"

	echo_session "github.com/labstack/echo-contrib/session"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Asdgo struct {
	*echo.Echo
}

type Config struct {
	Database *gorm.DB

	SessionName string

	CsrfName    string
	CsrfSkipper func(c echo.Context) bool

	TemplateNotFound templ.Component
}

func Env() {
	godotenv.Load()
}

func New(config *Config) *Asdgo {
	session.New(config.SessionName)
	validate.New()
	hash.New()
	mail.New()

	e := echo.New()
	e.HideBanner = true

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.Logger().Error(err)

		httpError, ok := err.(*echo.HTTPError)
		if ok {
			errorCode := httpError.Code
			switch errorCode {
			case http.StatusNotFound:
				template.Render(c, http.StatusNotFound, config.TemplateNotFound)
			default:
				// TODO: handle any other cases
			}
		}
	}

	e.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/static")
		},
		Format: "method=${method}, uri=${uri}, status=${status} (${remote_ip})\n",
	}))

	if config.CsrfName == "" {
		config.CsrfName = "asdgo_csrf"
	}
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: fmt.Sprintf("cookie:%s", config.CsrfName),
		Skipper: func(c echo.Context) bool {
			if config.CsrfSkipper == nil {
				return false
			}

			return config.CsrfSkipper(c)
		},
		CookieName:     config.CsrfName,
		CookiePath:     "/",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteLaxMode,
	}))

	e.Use(echo_session.Middleware(session.Instance.Store()))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := session.Instance.Get(c, "user_id")

			c.Set("userID", userID)

			// TODO: Figure out how todo this better...
			c.SetRequest(
				c.Request().WithContext(
					context.WithValue(c.Request().Context(), ctx.UserIDKey, userID),
				),
			)

			return next(c)
		}
	})

	if config.Database != nil {
		database.New(config.Database)
	}

	return &Asdgo{
		Echo: e,
	}
}

func Db() *database.Database {
	return database.Instance
}

func Session() *session.Session {
	return session.Instance
}

func Hash() *hash.Hash {
	return hash.Instance
}

func Validator() *validate.Validator {
	return validate.Instance
}

func Mailer() *mail.Mailer {
	return mail.Instance
}
