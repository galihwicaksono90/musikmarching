package middlewares

import (
	"encoding/json"
	"galihwicaksono90/musikmarching-be/internal/constants/model"

	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"net/http"
)

func ContributorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getSessionUser(r)

		if user.RoleName != db.RolenameContributor {
			response := model.Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
			json.NewEncoder(w).Encode(response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
