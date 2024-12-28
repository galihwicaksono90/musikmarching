package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	SessionName = "user-session"
)

type SessionOptions struct {
	CookiesKey string
	MaxAge     int
	HttpOnly   bool
	Secure     bool
}

func NewSessionStore(opts SessionOptions) *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(opts.CookiesKey))

	store.MaxAge(opts.MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = opts.HttpOnly
	store.Options.Secure = opts.Secure
	store.Options.SameSite = http.SameSiteLaxMode
	store.Options.Domain = "musikmarching.com"
	// store.Options.Secure = opts.Secure
	// store.Options.SameSite = http.SameSiteNoneMode

	return store
}
