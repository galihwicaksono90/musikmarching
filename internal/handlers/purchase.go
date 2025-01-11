package handlers

import (
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) HandleGetPurchasesByAccountID(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	purchases, err := h.purchase.GetPurchasesByAccountID(user.ID)

	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, "Purchases not found", purchases)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), purchases)
}

func (h *Handler) HandleGetPurchaseByID(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	purchaseID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, "Not a valid id", err)
		return
	}

	params := db.GetPurchaseByIdParams{
		ID:        purchaseID,
		AccountID: user.ID,
	}

	h.logger.Println(params, purchaseID)

	purchase, err := h.purchase.GetPurchaseByID(params)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, "Purchase not found", err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), purchase)
}

func (h *Handler) HandlePurchaseScore(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	id := mux.Vars(r)["id"]
	scoreId, err := uuid.Parse(id)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, "Score not found", err)
		return
	}

	purchaseID, err := h.purchase.PurchaseScore(user.ID, scoreId)
	if err != nil {
		h.handleResponse(w, http.StatusCreated, http.StatusText(http.StatusCreated), err)
		return
	}

	h.email.SendPurchaseInvoice(user)

	h.handleResponse(w, http.StatusCreated, http.StatusText(http.StatusCreated), purchaseID)
}

func (h *Handler) HandleUploadPaymentProof(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		// http.Error(w, "Failed to parse form", http.StatusBadRequest)
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	paymentID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	paymentProofUrl, err := h.file.UploadPaymentProof(r, "image_file")

	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	if err := h.purchase.UpdatePurchaseProof(db.UpdatePurchaseProofParams{
		PaymentProofUrl: pgtype.Text{
			String: paymentProofUrl,
			Valid:  true,
		},
		ID:        paymentID,
		AccountID: user.ID,
	}); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), true)
}

func (h *Handler) HandleGetPurchasedScoreById(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	id := mux.Vars(r)["id"]
	purchaseID, err := uuid.Parse(id)
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	purchase, err := h.purchase.GetPurchasedScoreById(purchaseID, user.ID)

	if err != nil {
		h.handleResponse(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), purchase)
}
