package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"

	"github.com/gorilla/mux"
)

func ScoreRouting(handler *handlers.Handler, router *mux.Router) {
	scoreRouter := router.PathPrefix("/score").Subrouter()

	scoreRouter.HandleFunc("/test", handler.HandleTestForm).Methods("POST")

	scoreRouter.Use(middlewares.AuthMiddleware, middlewares.ContributorMiddleware)
	scoreRouter.HandleFunc("/verified", handler.HandleGetScoresByCotributorId).Methods("GET")
	scoreRouter.HandleFunc("/create", handler.HandleCreateScore).Methods("POST")
	scoreRouter.HandleFunc("/update/{id}", handler.HandleUpdateScore).Methods("PUT")
	scoreRouter.HandleFunc("/verify/{id}", handler.HandleVerifyScore).Methods("POST")
}
