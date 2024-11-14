package handlers

import (
	"fmt"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"
	"strconv"
)

func (h *Handler) HandleGetVerifiedScores(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.FormValue("page_limit"))
	if err != nil {
		limit = 1
	}
	offset, err := strconv.Atoi(r.FormValue("page_offset"))
	if err != nil {
		offset = 0
	}

	scores := h.score.GetVerifiedScores(db.GetVerifiedScoresParams{
		Pagelimit:  int32(limit),
		Pageoffset: int32(offset),
	})

	verifiedScores := make([]components.VerifiedScoreProps, len(*scores))
	h.logger.Println(verifiedScores)

	for index, score := range *scores {
		s, _ := score.Price.Float64Value()
		fmt.Println(s)
		ss := fmt.Sprintf("%v", s.Float64)

		verifiedScores[index] = components.VerifiedScoreProps{
			ID:    score.ID.String(),
			Title: score.Title,
			Name:  score.Name,
			Price: ss,
		}
	}

	components.VerifiedScores(verifiedScores).Render(r.Context(), w)
}
