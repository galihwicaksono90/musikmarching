package handlers

import (
	"encoding/json"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	"galihwicaksono90/musikmarching-be/internal/services/account"
	"galihwicaksono90/musikmarching-be/internal/services/auth"
	"galihwicaksono90/musikmarching-be/internal/services/file"
	"galihwicaksono90/musikmarching-be/internal/services/purchase"
	"galihwicaksono90/musikmarching-be/internal/services/score"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/pkg/email"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger      *logrus.Logger
	store       *db.Store
	auth        auth.AuthService
	account     account.AccountService
	score       score.ScoreService
	purchase    purchase.PurchaseService
	file        file.FileService
	email       email.Email
	validate    *validator.Validate
}

func New(
	logger *logrus.Logger,
	store *db.Store,
	auth auth.AuthService,
	account account.AccountService,
	score score.ScoreService,
	purchase purchase.PurchaseService,
	file file.FileService,
	email email.Email,
	validate *validator.Validate,
) *Handler {
	return &Handler{
		logger,
		store,
		auth,
		account,
		score,
		purchase,
		file,
		email,
		validate,
	}
}

func (h *Handler) handleResponse(w http.ResponseWriter, code uint, message string, data interface{}) {
	response := model.Response(code, message, data)
	json.NewEncoder(w).Encode(response)
	h.logger.Info(data)
}
