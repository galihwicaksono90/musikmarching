package handlers

import (
	"encoding/json"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"
	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"
)

func (h *Handler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	session := middlewares.GetSession(r)

	if session == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		return
	}

	profile, err := h.profile.GetProfileById(session.ID)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	components.Profile(*profile).Render(r.Context(), w)
}
