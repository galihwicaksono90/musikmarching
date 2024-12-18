package handlers

import (
	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) HandlePurchaseScore(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)
	
	scoreID, err := uuid.Parse(mux.Vars(r)["score_id"])
	if err != nil {
		h.logger.Errorln(err)
		return
	}

	purchaseID, err := h.purchase.PurchaseScore(user, scoreID)
	if err != nil {
		h.logger.Errorln(err)
		return
	}

	if err != nil {
		h.logger.Errorln(err)
		return
	}

	h.email.SendPurchaseInvoice(user)

	h.logger.Printf("Purchase Score: %s", purchaseID)
}

func (h *Handler) HandleGetPurchases(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)

	purchases, err := h.purchase.GetPurchases(user.ID)

	if err != nil {
		components.Success("error").Render(r.Context(), w)
		return
	}

	components.PurchaseList(purchases).Render(r.Context(), w)
}
