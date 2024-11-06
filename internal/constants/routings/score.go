package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"

	"github.com/gorilla/mux"
)

func ScoreRouting(handler *handlers.Handler, router *mux.Router) {
	scoreRouter := router.PathPrefix("/score").Subrouter()

	scoreRouter.HandleFunc("/verified", handler.HandleGetVerifiedScores).Methods("GET")
}
