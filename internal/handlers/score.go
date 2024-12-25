package handlers

import (
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/utils"

	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) HandleGetScores(w http.ResponseWriter, r *http.Request) {
	scores, err := h.score.GetAll()
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "heyoooooo")
		return
	}
	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), scores)
}

func (h *Handler) HandleGetAllPublicScores(w http.ResponseWriter, r *http.Request) {
	scores, err := h.score.GetAllPublic()
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "heyoooooo")
		return
	}
	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), scores)
}

// func (h *Handler) HandleGetScores(w http.ResponseWriter, r *http.Request) {
// 	limit, offset := utils.ParsePagination(r)
//
// 	scores, err := h.score.GetAll(db.GetScoresPaginatedParams{
// 		Limit:  limit,
// 		Offset: offset,
// 	})
// 	if err != nil {
// 		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
// 		return
// 	}
//
// 	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), scores)
// }

func (h *Handler) HandleGetScoreById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	scoreId, err := uuid.Parse(id)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, "Score not found 1", err)
		return
	}

	score, err := h.score.GetById(scoreId)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, "Score not found 2", err)
		return
	}
	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), score)
}

func (h *Handler) HandleGetVerifiedScores(w http.ResponseWriter, r *http.Request) {
	pageLimit, pageOffset := utils.ParsePagination(r)

	scores, err := h.score.GetManyVerified(db.GetVerifiedScoresParams{
		PageLimit:  pageLimit,
		PageOffset: pageOffset,
	})
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), scores)
}
