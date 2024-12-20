package handlers

import (
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/utils"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) HandleGetContributorScores(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	limit, offset := utils.ParsePagination(r)

	scores, err := h.score.GetManyByContirbutorID(db.GetScoresByContributorIDParams{
		ID:         user.ID,
		Pageoffset: offset,
		Pagelimit:  limit,
	})

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

func (h *Handler) HandleGetContributorScore(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	scoreID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	score, err := h.score.GetOneByContributorID(db.GetScoreByContributorIDParams{
		ScoreID:       scoreID,
		ContributorID: user.ID,
	})

	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), score)
}

func (h *Handler) HandleCreateContributorScore(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	pdfUrl, images, err := h.file.UploadPdfFile(r, "pdf_file")

	for index, image := range images {
		h.logger.Infoln(image + strconv.Itoa(index))
	}

	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	audioUrl, err := h.file.UploadAudioFile(r, "audio_file")
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	title := r.FormValue("title")
	price, ok := utils.StringToBigInt(r.FormValue("price"))
	if !ok {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "invalid price")
		return
	}

	score, err := h.score.Create(model.CreateScoreDTO{
		ContributorID: user.ID,
		Title:         title,
		Price:         price,
		PdfUrl:        pdfUrl,
		PdfImageUrls:  images,
		AudioUrl:      audioUrl,
	})

	if err != nil {
		h.handleResponse(
			w,
			http.StatusInternalServerError,
			"Failed to create score",
			err,
		)
		return
	}

	h.handleResponse(w, http.StatusCreated, http.StatusText(http.StatusCreated), score)
}

func (h *Handler) HandleTesting(w http.ResponseWriter, r *http.Request) {
	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), "hello")
}

func (h *Handler) HandleUpdateContributorScore(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	scoreID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	price, ok := utils.StringToBigInt(r.FormValue("price"))
	if !ok {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "invalid price")
		return
	}

	params := model.UpdateScoreDTO{
		ContributorID: user.ID,
		Title: pgtype.Text{
			String: r.FormValue("title"),
			Valid:  true,
		},
		Price: pgtype.Numeric{
			Int:   price,
			Valid: true,
		},
	}

	pdfUrl, images, err := h.file.UploadPdfFile(r, "pdf_file")
	if pdfUrl != "" {
		params.PdfUrl = pgtype.Text{
			String: pdfUrl,
			Valid:  true,
		}
		params.PdfImageUrls = images
	}

	audioUrl, _ := h.file.UploadAudioFile(r, "audio_file")
	if audioUrl != "" {
		params.AudioUrl = pgtype.Text{
			String: audioUrl,
			Valid:  true,
		}
	}

	if err := h.score.Update(scoreID, params); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), false)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), user)
}
