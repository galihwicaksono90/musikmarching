package handlers

import (
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	user, _ := h.auth.GetSessionUser(r)

	scoreID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, "Not a valid id", err)
		return
	}

	purchase, err := h.purchase.GetPurchaseByID(db.GetPurchaseByIdParams{
		ScoreID:   scoreID,
		AccountID: user.ID,
	})

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
		h.handleResponse(w, http.StatusInternalServerError, "Score not found3", err)
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
