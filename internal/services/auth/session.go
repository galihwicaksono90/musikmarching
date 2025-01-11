package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionOptions struct {
	CookiesKey string
	MaxAge     int
	HttpOnly   bool
	Secure     bool
	Domain     string
	SameSite   http.SameSite
}

func NewSessionStore(opts SessionOptions) *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(opts.CookiesKey))

	// store.MaxAge(opts.MaxAge)
	store.Options.MaxAge = opts.MaxAge
	store.Options.Path = "/"
	store.Options.HttpOnly = opts.HttpOnly
	store.Options.Secure = opts.Secure
	store.Options.SameSite = http.SameSiteLaxMode
	if opts.Domain != "" {
		store.Options.Domain = opts.Domain
	}

	return store
}

func NewFileStore(opts SessionOptions) *sessions.FilesystemStore {
	store := sessions.NewFilesystemStore("./store/", []byte(opts.CookiesKey))

	// store.MaxAge(opts.MaxAge)
	store.Options.MaxAge = opts.MaxAge
	store.Options.Path = "/"
	store.Options.HttpOnly = opts.HttpOnly
	store.Options.Secure = opts.Secure
	store.Options.SameSite = http.SameSiteLaxMode
	if opts.Domain != "" {
		store.Options.Domain = opts.Domain
	}

	return store
}
