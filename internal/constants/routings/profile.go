package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"

	"github.com/gorilla/mux"
)

func ProfileRouting(handler *handlers.Handler, router *mux.Router) {
	profileRouter := router.PathPrefix("/profile").Subrouter()

	profileRouter.Use(middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	profileRouter.HandleFunc("", handler.HandleProfile).Methods("GET")
}
