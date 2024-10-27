package handlers

import (

	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"
)

func (h *Handler) HandleCreateScore(w http.ResponseWriter, r *http.Request) {
	session, _ := h.auth.GetSessionUser(r)

	profile, _:= h.profile.GetProfileByAccountId(session.ID)
	err := h.score.CreateScore(db.CreateScoreParams{
		ID:    profile.ID,
		Title: r.FormValue("title"),
	})

	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	scores, _:= h.score.GetScoresByAccountId(session.ID)

	components.Scores(&scores).Render(r.Context(), w)
}
