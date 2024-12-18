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

	purchaseRouter := router.PathPrefix("/purchase").Subrouter()
	purchaseRouter.Use(middlewares.AuthMiddleware)
	purchaseRouter.HandleFunc("", handler.HandleGetPurchasesByAccountID).Methods("GET")
	purchaseRouter.HandleFunc("/{id}", handler.HandleGetPurchaseByID).Methods("GET")
	purchaseRouter.HandleFunc("/{id}", handler.HandlePurchaseScore).Methods("POST")

	contributorRouter := router.PathPrefix("/contributor").Subrouter()
	contributorRouter.Use(middlewares.AuthMiddleware, middlewares.ContributorMiddleware)
	contributorRouter.HandleFunc("/scores", handler.HandleGetContributorScores).Methods("GET")
	contributorRouter.HandleFunc("/score/{id}", handler.HandleGetContributorScore).Methods("GET")
	contributorRouter.HandleFunc("/score", handler.HandleCreateContributorScore).Methods("POST")
	contributorRouter.HandleFunc("/score/{id}", handler.HandleUpdateContributorScore).Methods("PUT")
}
