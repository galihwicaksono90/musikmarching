package routings

import (
	"encoding/json"
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func ScoreRouting(handler *handlers.Handler, router *mux.Router) {
	router.HandleFunc("/score", handler.HandleCreateScore).Methods("POST")
	router.HandleFunc("/scoreee", func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
	}).Methods("GET")
}
