package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"

	"github.com/gorilla/mux"
)

func ProfileRouting(handler *handlers.Handler, router *mux.Router) {
	router.HandleFunc("/profile", handler.HandleProfile).Methods("GET")
}
