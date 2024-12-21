package middlewares

import (
	"encoding/json"
	"galihwicaksono90/musikmarching-be/internal/constants/model"

	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getSessionUser(r)

		if user == nil {
			response := model.Response(http.StatusForbidden, http.StatusText(http.StatusForbidden), nil)
			json.NewEncoder(w).Encode(response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
