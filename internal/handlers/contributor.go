package handlers

import (
	// "encoding/json"
	"encoding/json"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) HandleGetContributorScores(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	scores, err := h.score.GetManyByContirbutorID(user.ID)

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

	pdfUrl, images, err := h.file.UploadPdfFile(r, "pdf_file", 2)

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
	description := r.FormValue("description")

	difficulty := db.Difficulty("")
	if err := difficulty.Scan(r.FormValue("difficulty")); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Invalid difficulty value")
		return
	}
	contentType := db.ContentType("")
	if err := contentType.Scan(r.FormValue("content_type")); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Invalid content type value")
		return
	}

	price, ok := utils.StringToBigInt(r.FormValue("price"))
	if !ok {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "invalid price")
		return
	}

	dto := model.CreateScoreDTO{
		ContributorID: user.ID,
		Title:         title,
		Description:   description,
		Price:         price,
		PdfUrl:        pdfUrl,
		PdfImageUrls:  images,
		AudioUrl:      audioUrl,
		Difficulty:    difficulty,
		ContentType:   contentType,
	}

	scoreId, err := h.score.Create(dto)
	if err != nil {
		h.handleResponse(
			w,
			http.StatusInternalServerError,
			"Failed to create score",
			err,
		)
		return
	}

	instruments := strings.Split(r.FormValue("instruments"), ",")
	for _, i := range instruments {
		serial, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		h.instrument.CreateScoreInstrument(db.CreateScoreInstrumentParams{
			ScoreID:      scoreId,
			InstrumentID: int32(serial),
		})
	}

	categories := strings.Split(r.FormValue("categories"), ",")
	for _, c := range categories {
		serial, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		h.category.CreateScoreCategory(db.CreateScoreCategoryParams{
			ScoreID:    scoreId,
			CategoryID: int32(serial),
		})
	}

	allocations := strings.Split(r.FormValue("allocations"), ",")
	for _, a := range allocations {
		serial, err := strconv.Atoi(a)
		if err != nil {
			continue
		}
		h.allocation.CreateScoreAllocation(db.CreateScoreAllocationParams{
			ScoreID:      scoreId,
			AllocationID: int32(serial),
		})
	}

	h.handleResponse(w, http.StatusCreated, http.StatusText(http.StatusCreated), scoreId)
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

	difficulty := db.NullDifficulty{}
	if err := difficulty.Scan(r.FormValue("difficulty")); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Invalid difficulty value")
		return
	}
	contentType := db.NullContentType{}
	if err := contentType.Scan(r.FormValue("content_type")); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Invalid content type value")
		return
	}

	params := model.UpdateScoreDTO{
		ContributorID: user.ID,
		Title: pgtype.Text{
			String: r.FormValue("title"),
			Valid:  true,
		},
		Description: pgtype.Text{
			String: r.FormValue("description"),
			Valid:  true,
		},
		Price: pgtype.Numeric{
			Int:   price,
			Valid: true,
		},
		Difficulty:  difficulty,
		ContentType: contentType,
	}

	pdfUrl, images, err := h.file.UploadPdfFile(r, "pdf_file", 2)
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

	instruments := strings.Split(r.FormValue("instruments"), ",")
	instrumentsIDs := make([]int32, len(instruments))
	for _, i := range instruments {
		serial, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		instrumentsIDs = append(instrumentsIDs, int32(serial))
	}
	h.instrument.UpsertManyScoreInstrument(scoreID, instrumentsIDs)

	categories := strings.Split(r.FormValue("categories"), ",")
	categoryIDs := make([]int32, len(categories))
	for _, i := range categories {
		serial, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		categoryIDs = append(categoryIDs, int32(serial))
	}
	h.category.UpsertManyScoreCategory(scoreID, categoryIDs)

	allocations := strings.Split(r.FormValue("allocations"), ",")
	allocationIDs := make([]int32, len(allocations))
	for _, i := range allocations {
		serial, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		allocationIDs = append(allocationIDs, int32(serial))
	}
	h.allocation.UpsertManyScoreAllocation(scoreID, allocationIDs)

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), user)
}

func (h *Handler) HandleDeleteContributorScore(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	scoreID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	if err := h.score.Delete(db.DeleteScoreParams{
		ID:            scoreID,
		ContributorID: user.ID,
	}); err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), true)
}

type HandleCreateContributorInput struct {
	FullName string `json:"full_name"`
}

func (h *Handler) HandleCreateContributor(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)
	var input HandleCreateContributorInput

	if user.RoleName == db.RolenameContributor {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "This user is already a contributor")
		return
	}

	if user.RoleName != db.RolenameUser {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "This user is not allowed")
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Failed to decode input body")
		return
	}

	params := db.CreateContributorParams{
		ID:       user.ID,
		FullName: input.FullName,
	}

	_, err = h.contributor.Create(params)
	if err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Failed to create contributor")
		return
	}

	u := &model.SessionUser{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		RoleName:   db.RolenameContributor,
		PictureUrl: user.PictureUrl,
	}

	if err := h.auth.StoreUserSession(w, r, u); err != nil {
		h.logger.Println("err", err)
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Failed to store user session")
		return
	}

	h.handleResponse(w, http.StatusCreated, http.StatusText(http.StatusCreated), true)
}

func (h *Handler) HandleGetContributorScoreStatistics(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	statistics, err := h.contributor.GetScoreStatistics(user.ID)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), statistics)
}

func (h *Handler) HandleGetContributorBestSellingScores(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	scores, err := h.contributor.GetBestSellingScores(user.ID)
	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), scores)
}

func (h *Handler) HandleGetContributorPaymentMethod(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	method, err := h.contributor.GetPaymentMethod(user.ID)
	if err != nil {
		h.handleResponse(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), method)
}

type HandleUpdateContributorPaymentMethodInput struct {
	Method        string `json:"method"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
}

func (h *Handler) HandleUpsertContributorPaymentMethod(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	var input HandleUpdateContributorPaymentMethodInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.handleResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	err := h.contributor.UpsertPaymentMethod(db.UpsertContributorPaymentMethodParams{
		Method:        input.Method,
		AccountNumber: input.AccountNumber,
		ContributorID: user.ID,
		BankName:      input.BankName,
	})

	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), true)
}

func (h *Handler) HandleGetContributorPayments(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	payments, err := h.contributor.GetPayments(user.ID)

	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), payments)
}

func (h *Handler) HandleGetContributorPaymentStatistics(w http.ResponseWriter, r *http.Request) {
	user := h.getSessionUser(r)

	data, err := h.contributor.GetPaymentStatistics(user.ID)

	if err != nil {
		h.handleResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), data)
}
