package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"

	"github.com/gorilla/mux"
)

func Routings(handler *handlers.Handler, baseRouter *mux.Router) {
	router := baseRouter.PathPrefix("/api/v1").Subrouter()

	scoreRouter := router.PathPrefix("/score").Subrouter()

	scoreRouter.HandleFunc("", handler.HandleGetScores).Methods("GET")
	scoreRouter.HandleFunc("/{id}", handler.HandleGetScoreById).Methods("GET")

	scoreRouter.Use(middlewares.AuthMiddleware)
	scoreRouter.HandleFunc("/purchase/{id}", handler.HandlePurchaseScore).Methods("POST")
}
