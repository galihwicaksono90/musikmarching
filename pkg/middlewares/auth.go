package middlewares

import (
	"context"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	"galihwicaksono90/musikmarching-be/internal/services/auth"

	"net/http"

	"github.com/markbates/goth/gothic"
)

func GetSession(r *http.Request) *model.SessionUser {
	session, ok := r.Context().Value("user").(*model.SessionUser)
	if !ok {
		return nil
	}

	return session
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), auth.SessionName, model.SessionUser{})
		req := r.WithContext(ctx)

		session, _ := gothic.Store.Get(r, auth.SessionName)

		u := session.Values["user"]

		if u != nil {
			sessionUser := &model.SessionUser{
				ID:       u.(model.SessionUser).ID,
				Email:    u.(model.SessionUser).Email,
				Name:     u.(model.SessionUser).Name,
				RoleName: u.(model.SessionUser).RoleName,
			}
			ctx = context.WithValue(r.Context(), "user", sessionUser)
			req = r.WithContext(ctx)
		}

		next.ServeHTTP(w, req)
	})
}
