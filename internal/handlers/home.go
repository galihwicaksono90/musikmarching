package handlers

import (
	"galihwicaksono90/musikmarching-be/views/pages"
	"net/http"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)

	pages.Home(user).Render(r.Context(), w)
}
