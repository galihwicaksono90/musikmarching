package middlewares

import (
	"context"
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"

	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/spf13/viper"
)

const (
	UserContextName = "user"
)

func getSessionUser(r *http.Request) *model.SessionUser {
	u := r.Context().Value(UserContextName)
	if u == nil {
		return nil
	}

	user, ok := u.(*model.SessionUser)
	if !ok {
		return nil
	}

	return user
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionName := viper.GetString("SESSION_NAME")

		ctx := context.WithValue(r.Context(), UserContextName, nil)
		req := r.WithContext(ctx)

		session, err := gothic.Store.Get(r, sessionName)

		if err != nil {
			fmt.Errorf("session error: %s", err)
		}

		u := session.Values["user"]

		if u != nil {
			sessionUser := &model.SessionUser{
				ID:       u.(model.SessionUser).ID,
				Email:    u.(model.SessionUser).Email,
				Name:     u.(model.SessionUser).Name,
				RoleName: u.(model.SessionUser).RoleName,
			}

			ctx = context.WithValue(r.Context(), UserContextName, sessionUser)
			req = r.WithContext(ctx)
		}

		next.ServeHTTP(w, req)
	})
}
