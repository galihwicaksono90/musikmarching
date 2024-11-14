package handlers

import (
	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) HandlePurchaseScore(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)
	
	h.logger.Println("########yy")
	h.logger.Println(user)
	h.logger.Println("########yy")


	scoreID, err := uuid.Parse(mux.Vars(r)["score_id"])
	if err != nil {
		h.logger.Errorln(err)
		return
	}

	h.logger.Println("****")
	h.logger.Println(user.ID)
	h.logger.Println(scoreID)
	h.logger.Println("****")

	purchaseID, err := h.purchase.PurchaseScore(user.ID, scoreID)
	if err != nil {
		h.logger.Errorln(err)
		return
	}

	h.logger.Println(purchaseID)
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
