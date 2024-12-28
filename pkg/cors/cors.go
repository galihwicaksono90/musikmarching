package cors

import (
	"github.com/gorilla/handlers"
	"net/http"
)

func NewCorsHandler() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
    handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowCredentials(),
	)
}
