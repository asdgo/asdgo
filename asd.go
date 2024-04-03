package asdgo

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/asdgo/asdgo/acontext"
	"github.com/asdgo/asdgo/adatabase"
	"github.com/asdgo/asdgo/ahash"
	"github.com/asdgo/asdgo/amail"
	"github.com/asdgo/asdgo/asession"
	"github.com/asdgo/asdgo/atemplate"
	"github.com/asdgo/asdgo/avalidate"

	"github.com/labstack/echo-contrib/session"

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
	asession.New(config.SessionName)
	avalidate.New()
	ahash.New()

	if os.Getenv("MAIL_HOST") != "" {
		amail.New()
	}

	e := echo.New()
	e.HideBanner = true

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.Logger().Error(err)

		httpError, ok := err.(*echo.HTTPError)
		if ok {
			errorCode := httpError.Code
			switch errorCode {
			case http.StatusNotFound:
				if config.TemplateNotFound == nil {
					c.String(http.StatusNotFound, "Not Found")
					return
				}

				atemplate.Render(c, http.StatusNotFound, config.TemplateNotFound)
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

	e.Use(session.Middleware(asession.Instance.Store()))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := asession.Instance.Get(c, "user_id")

			c.Set("userID", userID)

			// TODO: Figure out how todo this better...
			c.SetRequest(
				c.Request().WithContext(
					context.WithValue(c.Request().Context(), acontext.UserIDKey, userID),
				),
			)

			return next(c)
		}
	})

	if config.Database != nil {
		adatabase.New(config.Database)
	}

	return &Asdgo{
		Echo: e,
	}
}

func Db() *adatabase.Database {
	return adatabase.Instance
}

func Session() *asession.Session {
	return asession.Instance
}

func Hash() *ahash.Hash {
	return ahash.Instance
}

func Validator() *avalidate.Validator {
	return avalidate.Instance
}

func Mailer() *amail.Mailer {
	return amail.Instance
}
