package handlers

import (
	"encoding/json"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"net/http"
)

func (h *Handler) HandleCreateContributorApply(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	if user.RoleName != "user" {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "This user is not allowed")
		return
	}
	var params db.CreateContributorApplyParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Not a valid request body")
		return
	}

	data, err := h.account.ApplyContributor(params)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), data)
}
