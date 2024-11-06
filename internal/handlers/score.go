package handlers

import (
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"
)

func (h *Handler) HandleCreateScoreForm(w http.ResponseWriter, r *http.Request) {
	components.CreateScoreForm().Render(r.Context(), w)
}

func (h *Handler) HandleCreateScore(w http.ResponseWriter, r *http.Request) {
	session, _ := h.auth.GetSessionUser(r)

	err := h.score.CreateScore(db.CreateScoreParams{
		ID:    session.ID,
		Title: r.FormValue("title"),
	})
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	// scores, _ := h.score.GetScoresByContributorId(session.ID)
	//
	// components.Scores(&scores).Render(r.Context(), w)
}

func (h *Handler) HandleGetScoresByContributor(w http.ResponseWriter, r *http.Request) {
	session, _ := h.auth.GetSessionUser(r)

	_, err := h.score.GetScoresByContributorId(session.ID)
	if err != nil {
		h.logger.Error(err)
		// scores = []db.Score{}
	}

	// components.Scores(&scores).Render(r.Context(), w)
}

func (h *Handler) HandleGetVerifiedScores(w http.ResponseWriter, r *http.Request) {
	score, _ := h.score.GetVerifiedScores()

	h.logger.Info("verified scores")

	components.Scores(score).Render(r.Context(), w)
}
