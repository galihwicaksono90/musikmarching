package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"

	"github.com/gorilla/mux"
)

func PageRouting(handler *handlers.Handler, router *mux.Router) {
	router.HandleFunc("/", handler.HandleHomePage).Methods("GET")

	contributorRouter := router.PathPrefix("/contributor").Subrouter()
	contributorRouter.Use(middlewares.AuthMiddleware, middlewares.ContributorMiddleware)
	contributorRouter.HandleFunc("", handler.HandleContributorPage).Methods("GET")
	contributorRouter.HandleFunc("/score/create", handler.HandleScoreCreatePage).Methods("GET")
	contributorRouter.HandleFunc("/score/update/{id}", handler.HandleScoreUpdatePage).Methods("GET")
}
