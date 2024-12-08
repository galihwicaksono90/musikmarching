package handlers

import (
	"encoding/json"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/utils"

	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) HandleGetScores(w http.ResponseWriter, r *http.Request) {
	limit, offset := utils.ParsePagination(r)

	scores, err := h.score.GetAll(db.GetScoresPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), scores)
}

func (h *Handler) HandleGetScoreById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	scoreId, err := uuid.Parse(id)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, "Score not found", err)
		return
	}

	score, err := h.score.GetById(scoreId)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, "Score not found", err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), score)
}
