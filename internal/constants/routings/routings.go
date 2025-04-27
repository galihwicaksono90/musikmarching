package routings

import (
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"

	"github.com/gorilla/mux"
)

func Routings(handler *handlers.Handler, baseRouter *mux.Router) {
	router := baseRouter.PathPrefix("/api/v1").Subrouter()

	router.Use(middlewares.SessionMiddleware)

	router.HandleFunc("/email", handler.HandleSendEmail).Methods("POST")
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/me", handler.HandleMe).Methods("GET")

	scoreRouter := router.PathPrefix("/score").Subrouter()
	scoreRouter.HandleFunc("", handler.HandleGetAllPublicScores).Methods("GET")
	scoreRouter.HandleFunc("/tags", handler.HandleGetScoreTags).Methods("GET")
	scoreRouter.HandleFunc("/{id}", handler.HandleGetPublicScoreById).Methods("GET")

	accountRouter := router.PathPrefix("/account").Subrouter()
	accountRouter.Use(middlewares.AuthMiddleware)
	accountRouter.HandleFunc("/contributor-request", handler.HandleCreateContributor).Methods("POST")

	purchaseRouter := router.PathPrefix("/purchase").Subrouter()
	purchaseRouter.Use(middlewares.AuthMiddleware)
	purchaseRouter.HandleFunc("", handler.HandleGetPurchasesByAccountID).Methods("GET")
	purchaseRouter.HandleFunc("/{id}", handler.HandleGetPurchaseByID).Methods("GET")
	purchaseRouter.HandleFunc("/{id}", handler.HandlePurchaseScore).Methods("POST")
	purchaseRouter.HandleFunc("/upload-proof/{id}", handler.HandleUploadPaymentProof).Methods("PUT")
	purchaseRouter.HandleFunc("/score/{id}", handler.HandleGetPurchasedScoreById).Methods("GET")
	purchaseRouter.HandleFunc("/scores/library", handler.HandleGetScoreLibrary).Methods("GET")

	contributorRouter := router.PathPrefix("/contributor").Subrouter()
	contributorRouter.Use(middlewares.AuthMiddleware, middlewares.ContributorMiddleware)
	contributorRouter.HandleFunc("/scores", handler.HandleGetContributorScores).Methods("GET")
	contributorRouter.HandleFunc("/score/{id}", handler.HandleGetContributorScore).Methods("GET")
	contributorRouter.HandleFunc("/score", handler.HandleCreateContributorScore).Methods("POST")
	contributorRouter.HandleFunc("/score/{id}", handler.HandleUpdateContributorScore).Methods("PUT")
	contributorRouter.HandleFunc("/scores/statistics", handler.HandleGetContributorScoreStatistics).Methods("GET")
	contributorRouter.HandleFunc("/scores/best-selling", handler.HandleGetContributorBestSellingScores).Methods("GET")
	contributorRouter.HandleFunc("/payment-method", handler.HandleGetContributorPaymentMethod).Methods("GET")
	contributorRouter.HandleFunc("/payment-method", handler.HandleUpsertContributorPaymentMethod).Methods("POST")
	contributorRouter.HandleFunc("/payments", handler.HandleGetContributorPayments).Methods("GET")
	contributorRouter.HandleFunc("/payments/statistics", handler.HandleGetContributorPaymentStatistics).Methods("GET")

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	adminRouter.HandleFunc("/scores", handler.HandleAdminGetScores).Methods("GET")
	adminRouter.HandleFunc("/score/verify/{id}", handler.HandleAdminVerifyScore).Methods("POST")
	adminRouter.HandleFunc("/contributors", handler.HandleAdminGetContributors).Methods("GET")
	adminRouter.HandleFunc("/contributor/verify/{id}", handler.HandleAdminVerifyContributor).Methods("POST")
	adminRouter.HandleFunc("/purchases", handler.HandleAdminGetPurchases).Methods("GET")
	adminRouter.HandleFunc("/purchase/verify/{id}", handler.HandleAdminVerifyPurchase).Methods("POST")
}
