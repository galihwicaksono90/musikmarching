package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"

	"github.com/gorilla/mux"
)

func AuthRouting(handler *handlers.Handler, router *mux.Router) {
	r := router.PathPrefix("/oauth2").Subrouter()
	r.HandleFunc("/{provider}", handler.HandleProviderLogin).Methods("GET")
	r.HandleFunc("/{provider}/callback", handler.HandleAuthCallbackFunction).Methods("GET")
	r.HandleFunc("/logout/{provider}", handler.HandleLogout).Methods("GET")
}
