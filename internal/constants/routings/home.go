package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"

	"github.com/gorilla/mux"
)

func HomeRouting(handler *handlers.Handler, router *mux.Router) {
	router.HandleFunc("/", handler.HandleHome).Methods("GET")
}
