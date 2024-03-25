package session

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var Instance *sessions.CookieStore

func New() {
	store := sessions.NewCookieStore([]byte(os.Getenv("APP_KEY")))
	store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   86400 * 7,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}

	Instance = store
}

func Get(r *http.Request, key string) string {
	sess, _ := Instance.Get(r, "checkeroni_session")

	if _, ok := sess.Values[key]; !ok {
		return ""
	}

	if ok := sess.Values[key].(string); ok != "" {
		return ok
	}

	return ""
}

func Set(w http.ResponseWriter, r *http.Request, key string, value string) {
	sess, _ := Instance.Get(r, "checkeroni_session")
	sess.Values[key] = value

	err := sess.Save(r, w)
	if err != nil {
		fmt.Printf("[SESSION ERROR]: %v\n", err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request, key string) {
	sess, _ := Instance.Get(r, "checkeroni_session")
	sess.Options.MaxAge = -1

	err := sess.Save(r, w)
	if err != nil {
		fmt.Printf("[SESSION ERROR]: %v\n", err)
	}
}
