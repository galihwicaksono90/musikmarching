package handlers

import (
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/views/components"
	"galihwicaksono90/musikmarching-be/views/pages"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) HandleHomePage(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)
	if user.RoleName == db.RolenameAdmin {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	scores, err := h.score.GetAll()
	if err != nil {
		h.logger.Errorln(err)
		return
	}

	pages.HomePage(user, scores).Render(r.Context(), w)
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

func (h *Handler) HandleAdminPage(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)

	pages.AdminPage(user).Render(r.Context(), w)
}

func (h *Handler) HandleAdminScoresPage(w http.ResponseWriter, r *http.Request) {
	scores, err := h.score.GetAll()
	if err != nil {
		h.logger.Errorln(err)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	for _, score := range scores {
		h.logger.Println(score)
	}

	pages.AdminScoresPage(scores).Render(r.Context(), w)
}

func (h *Handler) HandleAdminScorePage(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	scoreId, err := uuid.Parse(id)
	if err != nil {
		h.logger.Errorln(err)
		http.Redirect(w, r, "/admin/scores", http.StatusSeeOther)
		return
	}

	score, err := h.score.GetById(scoreId)
	if err != nil {
		h.logger.Errorln(err)
		http.Redirect(w, r, "/admin/scores", http.StatusSeeOther)
		return
	}

	pages.AdminScorePage(score).Render(r.Context(), w)
}
