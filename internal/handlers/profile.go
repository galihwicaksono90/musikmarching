package handlers

import (
	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"
)

func (h *Handler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	session, err := h.auth.GetSessionUser(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}
	h.logger.Println(session)

	profile, err := h.profile.GetProfileByAccountId(session.ID)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	components.Profile(*profile).Render(r.Context(), w)
}
