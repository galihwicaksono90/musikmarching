package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"

	"github.com/gorilla/mux"
)

func ScoreRouting(handler *handlers.Handler, router *mux.Router) {
	scoreRouter := router.PathPrefix("/score").Subrouter()

	scoreRouter.HandleFunc("/verified", handler.HandleGetVerifiedScores).Methods("GET")
	scoreRouter.HandleFunc("/create", handler.HandleCreateScore).Methods("POST")
	scoreRouter.HandleFunc("/update/{id}", handler.HandleUpdateScore).Methods("PUT")
}
