package handlers

import (
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/utils"

	"galihwicaksono90/musikmarching-be/views/components"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"

	_ "github.com/go-playground/validator/v10"
)

func (h *Handler) HandleGetScoresByCotributorId(w http.ResponseWriter, r *http.Request) {
	user, _ := h.auth.GetSessionUser(r)

	limit, offset := utils.ParsePagination(r)

	scores, err := h.score.GetByContirbutorID(db.GetScoresByContributorIDParams{
		ID:         user.ID,
		Pagelimit:  limit,
		Pageoffset: offset,
	})
	if err != nil {
		h.logger.Errorln(err)
		return
	}

	verifiedScores := make([]components.VerifiedScoreProps, len(scores))
	h.logger.Println(verifiedScores)

	for index, score := range scores {
		s, _ := score.Price.Float64Value()
		ss := fmt.Sprintf("%v", s.Float64)

		verifiedScores[index] = components.VerifiedScoreProps{
			ID:         score.ID.String(),
			Title:      score.Title,
			Name:       score.Name,
			Price:      ss,
			IsVerified: score.IsVerified,
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
			Int:   price,
			Valid: true,
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

func (h *Handler) HandleVerifyScore(w http.ResponseWriter, r *http.Request) {
	scoreID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Errorln(err)
		return
	}

	if err := h.score.Verify(scoreID); err != nil {
		h.logger.Errorln(err)
		return
	}

	limit, offset := utils.ParsePagination(r)

	scores, err := h.score.GetAll(db.GetScoresPaginatedParams{
		Limit:   limit,
		Offset:  offset,
		Column3: nil,
	})

	if err != nil {
		h.logger.Errorln(err)
		return
	}

	components.AdminScoreList(scores).Render(r.Context(), w)
}

func validateForm(r *http.Request) map[string]string {
	errors := make(map[string]string)

	title := r.FormValue("titlee")
	if title == "" {
		errors["title"] = "Title is required"
	} else if len(title) < 5 {
		errors["title"] = "Length must be greater than 5"
	} else if len(title) > 100 {
		errors["title"] = "Length must be less than 100"
	}

	return errors
}

func (h *Handler) HandleTestForm(w http.ResponseWriter, r *http.Request) {
	errors := validateForm(r)
	values := make(map[string]string)

	for k, v := range r.Form {
		values[k] = v[0]
	}

	h.logger.Println("errors======")
	h.logger.Println(errors)
	h.logger.Println("errors======")

	components.TestForm(values, errors).Render(r.Context(), w)
}
