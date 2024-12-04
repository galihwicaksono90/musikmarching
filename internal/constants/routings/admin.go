package routings
import (
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"

	"github.com/gorilla/mux"
)

func AdminRouting(handler *handlers.Handler, router *mux.Router) {
	adminRouter := router.PathPrefix("/admin").Subrouter()

	adminRouter.Use(middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	adminRouter.HandleFunc("/score/verify/{id}", handler.HandleVerifyScore).Methods("POST")
}
