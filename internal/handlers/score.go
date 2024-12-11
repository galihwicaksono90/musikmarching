package handlers

import (
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

func (h *Handler) HandlePurchaseScore(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)

	scoreID, err := uuid.Parse(mux.Vars(r)["score_id"])
	if err != nil {
		h.logger.Errorln(err)
		return
	}

	purchaseID, err := h.purchase.PurchaseScore(user, scoreID)
	if err != nil {
		h.logger.Errorln(err)
		return
	}

	if err != nil {
		h.logger.Errorln(err)
		return
	}

	h.email.SendPurchaseInvoice(user)

	h.logger.Printf("Purchase Score: %s", purchaseID)
}
