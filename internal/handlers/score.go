package handlers

import (
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/utils"

	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
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

	scores := h.score.GetVerified(db.GetVerifiedScoresParams{
		Pagelimit:  int32(limit),
		Pageoffset: int32(offset),
	})

	verifiedScores := make([]components.VerifiedScoreProps, len(*scores))
	h.logger.Println(verifiedScores)

	for index, score := range *scores {
		s, _ := score.Price.Float64Value()
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

func (h *Handler) HandleCreateScore(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetSessionUser(r)
	if err != nil {
		h.logger.Errorln(err)
		components.Success(err.Error()).Render(r.Context(), w)
		return
	}

	price, ok := utils.StringToBigInt(r.FormValue("price"))
	if !ok {
		components.Success("failed to parse price").Render(r.Context(), w)
		return
	}

	// upload pdf
	h.logger.Println("Uploading PDF")
	pdfUploadUrl, err := h.score.UploadPdfFile(r)
	if err != nil {
		h.logger.Errorln(err)
		components.Success(err.Error()).Render(r.Context(), w)
		return
	}

	// upload music
	h.logger.Println("Uploading Music")
	musicUploadUrl, err := h.score.UploadMusicFile(r)
	if err != nil {
		h.logger.Errorln(err)
		components.Success(err.Error()).Render(r.Context(), w)
	}

	if _, err := h.score.Create(model.CreateScoreDTO{
		ContributorID: user.ID,
		Title:         r.FormValue("title"),
		Price:         price,
		PdfUrl:        pdfUploadUrl,
		MusicUrl:      musicUploadUrl,
	}); err != nil {
		h.logger.Errorln(err)
		components.Success(err.Error()).Render(r.Context(), w)
		return
	}

	hxRedirectWithToast(w, r, "/contributor", "Score created successfully")
}

func (h *Handler) HandleUpdateScore(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)

	scoreId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Errorln(err)
		components.Success(err.Error()).Render(r.Context(), w)
		return
	}

	params := model.UpdateScoreDTO{
		ContributorID: user.ID,
	}

	if r.FormValue("title") != "" {
		params.Title = pgtype.Text{
			String: r.FormValue("title"),
			Valid:  true,
		}
	}

	if r.FormValue("price") != "" {
		price, ok := utils.StringToBigInt(r.FormValue("price"))
		if !ok {
			components.Success("failed to parse price").Render(r.Context(), w)
			return
		}

		params.Price = pgtype.Numeric{
			Int: price,
			Valid:  true,
		}
	}

	err = h.score.Update(scoreId, params)

	if err != nil {
		h.logger.Errorln(err)
		components.Success(err.Error()).Render(r.Context(), w)
		return
	}

	components.Success("Score updated successfully").Render(r.Context(), w)
}
