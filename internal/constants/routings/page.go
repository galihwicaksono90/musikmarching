package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"

	"github.com/gorilla/mux"
)

func PageRouting(handler *handlers.Handler, router *mux.Router) {
	router.HandleFunc("/", handler.HandleHomePage).Methods("GET")

	contributorRouter := router.PathPrefix("/").Subrouter()
	contributorRouter.Use(middlewares.AuthMiddleware, middlewares.ContributorMiddleware)
	contributorRouter.HandleFunc("/contributor", handler.HandleContributorPage).Methods("GET")
}
