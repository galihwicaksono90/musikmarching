package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) HandleAdminGetScores(w http.ResponseWriter, r *http.Request) {
	scores, err := h.score.GetAll()

	if err != nil {
		h.handleResponse(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err,
		)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), scores)
}

func (h *Handler) HandleAdminVerifyScore(w http.ResponseWriter, r *http.Request) {
	scoreID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	err = h.score.Verify(scoreID)
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), nil)
}

func (h *Handler) HandleAdminGetContributors(w http.ResponseWriter, r *http.Request) {
	res, err := h.contributor.GetAll()
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}
	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (h *Handler) HandleAdminVerifyContributor(w http.ResponseWriter, r *http.Request) {
	scoreID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	err = h.contributor.Verify(scoreID)
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), true)
}

func (h *Handler) HandleAdminGetPurchases(w http.ResponseWriter, r *http.Request) {
	purchases, err := h.purchase.GetAll()
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}
	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), purchases)
}

func (h *Handler) HandleAdminVerifyPurchase(w http.ResponseWriter, r *http.Request) {
	purchaseID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	err = h.purchase.Verify(purchaseID)
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), true)
}
