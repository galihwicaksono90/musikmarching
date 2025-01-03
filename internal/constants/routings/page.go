package routings

import (
	"encoding/json"
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func PageRouting(handler *handlers.Handler, router *mux.Router) {
	router.HandleFunc("/svelte", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello from Goland!")
	})
	router.Use(middlewares.AuthMiddleware)
	router.HandleFunc("/svelte/me", handler.HandleMe).Methods("GET")

	// contributorRouter := router.PathPrefix("/contributor").Subrouter()
	// contributorRouter.Use(middlewares.AuthMiddleware, middlewares.ContributorMiddleware)
	// contributorRouter.HandleFunc("", handler.HandleContributorPage).Methods("GET")
	// contributorRouter.HandleFunc("/score/create", handler.HandleScoreCreatePage).Methods("GET")
	// contributorRouter.HandleFunc("/score/update/{id}", handler.HandleScoreUpdatePage).Methods("GET")
	//
	// adminRouter := router.PathPrefix("/admin").Subrouter()
	// adminRouter.Use(middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	// adminRouter.HandleFunc("", handler.HandleAdminPage).Methods("GET")
	// adminRouter.HandleFunc("/scores", handler.HandleAdminScoresPage).Methods("GET")
	// adminRouter.HandleFunc("/score/{id}", handler.HandleAdminScorePage).Methods("GET")
}
