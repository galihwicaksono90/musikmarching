package handlers

import (
	"encoding/json"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"net/http"
)

func (h *Handler) HandleGetContributorApplyByAccountID(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	contributorApply, err := h.contributorApply.GetByAccountID(user.ID)
	if err != nil {
		h.handleResponse(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), "This user is not applied")
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), contributorApply)
}

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

	_, err := h.contributorApply.GetByAccountID(user.ID)
	if err == nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "This user already applied")
		return
	}

	params.ID = user.ID
	data, err := h.contributorApply.Apply(params)
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	h.handleResponse(w, http.StatusCreated, http.StatusText(http.StatusOK), data)
}

func (h *Handler) HandleUpdateContributorApply(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("HandleUpdateContributorApply")
	user := h.getSessionUser(r)
	if user.RoleName == "contributor" {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "This user is not allowed")
	}

	var params db.UpdateContributorApplyParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Not a valid request body")
		return
	}

	_, err := h.contributorApply.GetByAccountID(user.ID)
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "User not found")
		return
	}

	params.AccountID = user.ID
	if err := h.contributorApply.Update(params); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), nil)
}

func (h *Handler) HandleGetContributorApplications(w http.ResponseWriter, r *http.Request) {
	applications, err := h.contributorApply.GetAll()
	if err != nil {
		h.logger.Error(err)
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), applications)
}
