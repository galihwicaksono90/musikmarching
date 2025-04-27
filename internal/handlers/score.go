package handlers

import (
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/utils"

	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) HandleGetScores(w http.ResponseWriter, r *http.Request) {
	scores, err := h.score.GetAll()
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}
	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), scores)
}

func (h *Handler) HandleGetAllPublicScores(w http.ResponseWriter, r *http.Request) {
	scores, err := h.score.GetAllPublic(r.URL.Query())
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	data := make([]db.ScorePublicView, len(scores))

	for i, s := range scores {
		data[i] = s.ScorePublicView
	}

	var count int64
	if len(scores) > 0 {
		count = scores[0].Count
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), model.PublicScores{
		Scores: data,
		Count:  count,
	})
}

func (h *Handler) HandleGetPublicScoreById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	scoreId, err := uuid.Parse(id)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, "Score not found", err)
		return
	}

	score, err := h.score.GetPublicById(scoreId)
	if err != nil {
		h.handleResponse(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), err)
		return
	}
	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), score)
}

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
	pageLimit, pageOffset := utils.ParsePagination(r.URL.Query())

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

func (h *Handler) HandleGetScoreLibrary(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	scores, err := h.score.GetScoreLibrary(user.ID, r.URL.Query())
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	data := make([]db.ScoreLibraryView, len(scores))

	for i, s := range scores {
		data[i] = s.ScoreLibraryView
	}

	var count int64
	if len(scores) > 0 {
		count = scores[0].Count
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), model.LibraryScores{
		Scores: data,
		Count:  count,
	})
}

type Result struct {
	Scores []db.ScorePublicView `json:"scores"`
	Count  int64                `json:"count"`
}
