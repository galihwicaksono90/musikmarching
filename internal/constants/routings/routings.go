package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"

	"github.com/gorilla/mux"
)

func Routings(handler *handlers.Handler, baseRouter *mux.Router) {
	router := baseRouter.PathPrefix("/api/v1").Subrouter()

	scoreRouter := router.PathPrefix("/score").Subrouter()

	scoreRouter.HandleFunc("", handler.HandleGetScores).Methods("GET")
	scoreRouter.HandleFunc("/{id}", handler.HandleGetScoreById).Methods("GET")
}
