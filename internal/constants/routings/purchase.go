package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"

	"github.com/gorilla/mux"
)

func PurchaseRouting(handler *handlers.Handler, router *mux.Router) {
	scoreRouter := router.PathPrefix("/purchase").Subrouter()

	scoreRouter.HandleFunc("/score/{score_id}", handler.HandlePurchaseScore).Methods("POST")
	scoreRouter.HandleFunc("/purchases", handler.HandleGetPurchases).Methods("GET")
}
