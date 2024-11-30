package middlewares

import (
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"net/http"
)

func ContributorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := GetSession(r)

		if session == nil || session.RoleName != db.RolenameContributor {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
