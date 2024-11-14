package handlers

import (
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/views/pages"
	"net/http"
)

func (h *Handler) HandleHomePage(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)
	purchases, _ := h.purchase.GetPurchases(user.ID)

	pages.Home(user, purchases).Render(r.Context(), w)
}

func (h *Handler) HandleContributorPage(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)
	if user.RoleName != db.RolenameContributor{
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}

	pages.Home(user, []db.Purchase{}).Render(r.Context(), w)
}
