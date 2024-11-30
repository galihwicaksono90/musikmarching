package handlers

import (
	"galihwicaksono90/musikmarching-be/internal/services/account"
	"galihwicaksono90/musikmarching-be/internal/services/auth"
	"galihwicaksono90/musikmarching-be/internal/services/purchase"
	"galihwicaksono90/musikmarching-be/internal/services/score"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/pkg/email"
	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"

	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

type Handler struct {
  logger *logrus.Logger
  store *db.Store
  auth auth.AuthService
  account account.AccountService
  score score.ScoreService
  purchase purchase.PurchaseService
  fileStorage *minio.Client
  email email.Email
}

func New(
  logger * logrus.Logger, 
  store *db.Store, 
  auth auth.AuthService, 
  account account.AccountService,
  score score.ScoreService,
  purchase purchase.PurchaseService,
  fileStorage *minio.Client,
  email email.Email,
) *Handler {
  return &Handler{
    logger,
    store,
    auth,
    account,
    score,
    purchase,
    fileStorage,
    email,
  }
}

func hxRedirect(w http.ResponseWriter, url string) {
  w.Header().Set("HX-Redirect", url)
  w.WriteHeader(http.StatusOK) // OK response
}

func hxRedirectWithToast(w http.ResponseWriter, r *http.Request, url string, message string) {
  w.Header().Set("HX-Redirect", url)
  w.WriteHeader(http.StatusOK) // OK response

  components.Success(message).Render(r.Context(), w)
}
