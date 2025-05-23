package handlers

import (
	"encoding/json"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	"galihwicaksono90/musikmarching-be/internal/services/account"
	"galihwicaksono90/musikmarching-be/internal/services/admin"
	"galihwicaksono90/musikmarching-be/internal/services/allocation"
	"galihwicaksono90/musikmarching-be/internal/services/auth"
	"galihwicaksono90/musikmarching-be/internal/services/category"
	"galihwicaksono90/musikmarching-be/internal/services/contributor"
	"galihwicaksono90/musikmarching-be/internal/services/contributor-apply"
	"galihwicaksono90/musikmarching-be/internal/services/file"
	"galihwicaksono90/musikmarching-be/internal/services/instrument"
	"galihwicaksono90/musikmarching-be/internal/services/payment"
	"galihwicaksono90/musikmarching-be/internal/services/purchase"
	"galihwicaksono90/musikmarching-be/internal/services/score"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/pkg/email"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger           *logrus.Logger
	store            *db.Store
	auth             auth.AuthService
	account          account.AccountService
	admin          admin.AdminService
	score            score.ScoreService
	purchase         purchase.PurchaseService
	payment          payment.PaymentService
	contributor      contributor.ContributorService
	contributorApply contributorapply.ContributorApplyService
	instrument       instrument.InstrumentService
	category         category.CategoryService
	allocation       allocation.AllocationService
	file             file.FileService
	email            email.Email
	validate         *validator.Validate
}

func New(
	logger *logrus.Logger,
	store *db.Store,
	auth auth.AuthService,
	account account.AccountService,
	admin admin.AdminService,
	score score.ScoreService,
	purchase purchase.PurchaseService,
	payment payment.PaymentService,
	contributor contributor.ContributorService,
	contributorApply contributorapply.ContributorApplyService,
	instrument instrument.InstrumentService,
	category category.CategoryService,
	allocation allocation.AllocationService,
	file file.FileService,
	email email.Email,
	validate *validator.Validate,
) *Handler {
	return &Handler{
		logger,
		store,
		auth,
		account,
		admin,
		score,
		purchase,
		payment,
		contributor,
		contributorApply,
		instrument,
		category,
		allocation,
		file,
		email,
		validate,
	}
}

func (h *Handler) handleResponse(w http.ResponseWriter, code uint, message string, data interface{}) {
	response := model.Response(code, message, data)
	json.NewEncoder(w).Encode(response)
	// h.logger.Info(data)
}
