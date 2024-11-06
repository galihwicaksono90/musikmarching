package middlewares

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"net/http"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := GetSession(r)

		if session == nil || session.RoleName != db.RolenameAdmin {
			ctx := context.WithValue(r.Context(), "user", nil)
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		}

		next.ServeHTTP(w, r)
	})
}
