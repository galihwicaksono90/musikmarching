package handlers

import (
	"galihwicaksono90/musikmarching-be/views/components"
	"galihwicaksono90/musikmarching-be/views/pages"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) HandleHomePage(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)
	purchases, _ := h.purchase.GetPurchases(user.ID)

	pages.HomePage(user, purchases).Render(r.Context(), w)
}

func (h *Handler) HandleContributorPage(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)

	pages.ContributorPage(user).Render(r.Context(), w)
}

func (h *Handler) HandleScoreCreatePage(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)

	pages.ScoreCreatePage(user).Render(r.Context(), w)
}

func (h *Handler) HandleScoreUpdatePage(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)
	id := mux.Vars(r)["id"] 

	scoreId, err := uuid.Parse(id)
	if err != nil {
		h.logger.Errorln(err)
		http.Redirect(w, r, "/contributor", http.StatusSeeOther)
		return
	}

	score, err := h.score.GetById(scoreId)
	if err != nil {
		h.logger.Errorln(err)
		http.Redirect(w, r, "/contributor", http.StatusSeeOther)
		return
	}

	pages.ScoreUpdatePage(user, id, components.ScoreFormProps{
		Title: score.Title,
		Price: score.Price.Int.String(),
	}).Render(r.Context(), w)
}
