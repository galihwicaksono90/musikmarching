package score

import (
	"context"
	"errors"
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	// "galihwicaksono90/musikmarching-be/internal/services/instrument"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/utils"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type ScoreService interface {
	GetAllPublic(url.Values) ([]db.GetAllPublicScoresRow, error)
	GetPublicById(uuid.UUID) (db.ScorePublicView, error)
	GetScoreLibrary(uuid.UUID, url.Values) ([]db.GetScoreLibraryRow, error)
	GetManyByContributorId(account_id uuid.UUID) ([]db.Score, error)
	Create(model.CreateScoreDTO) (uuid.UUID, error)
	Update(uuid.UUID, model.UpdateScoreDTO) error
	Delete(db.DeleteScoreParams) error
	GetManyVerified(db.GetVerifiedScoresParams) (*[]db.GetVerifiedScoresRow, error)
	GetVerifiedById(id uuid.UUID) (db.GetVerifiedScoreByIdRow, error)
	GetById(id uuid.UUID) (db.Score, error)
	GetManyByContirbutorID(uuid.UUID) ([]db.GetScoresByContributorIDRow, error)
	GetOneByContributorID(db.GetScoreByContributorIDParams) (db.ScoreContributorView, error)
	GetAll() ([]db.Score, error)
	Verify(id uuid.UUID) error
}

type scoreService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetScoreLibrary implements ScoreService.
func (s *scoreService) GetScoreLibrary(id uuid.UUID, urlValues url.Values) ([]db.GetScoreLibraryRow, error) {
	limit, offset := utils.ParsePagination(urlValues)
	ctx := context.Background()
	params := db.GetScoreLibraryParams{
		PageOffset: offset,
		PageLimit:  limit,
		AccountID:  id,
	}

	return s.store.GetScoreLibrary(ctx, params)
}

// GetPublicById implements ScoreService.
func (s *scoreService) GetPublicById(id uuid.UUID) (db.ScorePublicView, error) {
	ctx := context.Background()
	return s.store.GetPublicScoreById(ctx, id)
}

// GetAllPublicScores implements ScoreService.
func (s *scoreService) GetAllPublic(urlValues url.Values) ([]db.GetAllPublicScoresRow, error) {
	limit, offset := utils.ParsePagination(urlValues)
	title := urlValues.Get("title")
	title = fmt.Sprintf("%%%s%%", title)
	instruments := urlValues["instruments"]
	categories := urlValues["categories"]
	allocations := urlValues["allocations"]
	difficulty := urlValues.Get("difficulty")

	params := db.GetAllPublicScoresParams{}
	params.PageLimit = limit
	params.PageOffset = offset
	params.Title = title
	params.Instruments = instruments
	params.Categories = categories
	params.Allocations = allocations

	var difficultyFilter db.NullDifficulty
	if err := difficultyFilter.Scan(difficulty); err == nil && isValidDifficulty(string(difficultyFilter.Difficulty)) {
		params.Difficulty = difficultyFilter
	}

	contentType := db.NullContentType{}
	if err := contentType.Scan(urlValues.Get("content_type")); err == nil && isValidContentType(string(contentType.ContentType)) {
		params.ContentType = contentType
	}

	result, err := s.store.GetAllPublicScores(context.Background(), params)
	return result, err
}

// GetOneByContributorID implements ScoreService.
func (s *scoreService) GetOneByContributorID(params db.GetScoreByContributorIDParams) (db.ScoreContributorView, error) {
	return s.store.GetScoreByContributorID(context.Background(), params)
}

// GetByContirbutorID implements ScoreService.
func (s *scoreService) GetManyByContirbutorID(id uuid.UUID) ([]db.GetScoresByContributorIDRow, error) {
	s.logger.Println(id)
	ctx := context.Background()
	return s.store.GetScoresByContributorID(ctx, id)
}

// VerifyScore implements ScoreService.
func (s *scoreService) Verify(id uuid.UUID) error {
	return s.store.VerifyScore(context.Background(), id)
}

// GetAll implements ScoreService.
func (s *scoreService) GetAll() ([]db.Score, error) {
	ctx := context.Background()
	result, err := s.store.GetScoresPaginated(ctx)
	if err != nil {
		{
		}
		return nil, err
	}

	return result, nil
}

// GetScoreById implements ScoreService.
func (s *scoreService) GetById(id uuid.UUID) (db.Score, error) {
	ctx := context.Background()
	return s.store.GetScoreById(ctx, id)
}

// GetManyByContributorId implements ScoreService.
func (s *scoreService) GetManyByContributorId(account_id uuid.UUID) ([]db.Score, error) {
	panic("unimplemented")
}

// GetVerified implements ScoreService.
func (s *scoreService) GetManyVerified(params db.GetVerifiedScoresParams) (*[]db.GetVerifiedScoresRow, error) {
	scores, err := s.store.GetVerifiedScores(context.Background(), params)

	if err != nil {
		return &[]db.GetVerifiedScoresRow{}, err
	}

	return &scores, err
}

// GetVerifiedById implements ScoreService.
func (s *scoreService) GetVerifiedById(id uuid.UUID) (db.GetVerifiedScoreByIdRow, error) {
	return s.store.GetVerifiedScoreById(context.Background(), id)
}

// CreateScore implements ScoreService.
func (s *scoreService) Create(params model.CreateScoreDTO) (uuid.UUID, error) {
	ctx := context.Background()

	return s.store.CreateScore(ctx, db.CreateScoreParams{
		Title: params.Title,
		Description: pgtype.Text{
			String: params.Description,
			Valid:  true,
		},
		Price: pgtype.Numeric{
			Int:   params.Price,
			Valid: true,
		},
		PdfUrl:        params.PdfUrl,
		PdfImageUrls:  params.PdfImageUrls,
		AudioUrl:      params.AudioUrl,
		ContributorID: params.ContributorID,
		Difficulty:    params.Difficulty,
		ContentType:   params.ContentType,
	})
}

// Update implements ScoreService.
func (s *scoreService) Update(scoreId uuid.UUID, params model.UpdateScoreDTO) error {
	ctx := context.Background()

	scoreCheck, err := s.store.GetScoreById(ctx, scoreId)
	if err != nil {
		return err
	}

	if scoreCheck.ContributorID != params.ContributorID {
		return errors.New("You are not the owner of this score")
	}

	if err := s.store.UpdateScore(ctx, db.UpdateScoreParams{
		Title:        params.Title,
		Price:        params.Price,
		Description:  params.Description,
		Difficulty:   params.Difficulty,
		ContentType:  params.ContentType,
		PdfUrl:       params.PdfUrl,
		PdfImageUrls: params.PdfImageUrls,
		AudioUrl:     params.AudioUrl,
		ID:           scoreId,
	}); err != nil {
		return err
	}

	return nil
}

func (s *scoreService) Delete(params db.DeleteScoreParams) error {
	ctx := context.Background()

	return s.store.DeleteScore(ctx, params)
}

// GetScoresByContributorId implements ScoreService.
func (s *scoreService) GetScoresByContributorId(account_id uuid.UUID) ([]db.Score, error) {
	panic("unimplemented")
}

func isValidDifficulty(difficulty string) bool {
	switch db.Difficulty(difficulty) {
	case db.DifficultyBeginner, db.DifficultyAdvanced, db.DifficultyIntermediate:
		return true
	}
	return false
}

func isValidContentType(contentType string) bool {
	switch db.ContentType(contentType) {
	case db.ContentTypeExclusive, db.ContentTypeNonExclusive:
		return true
	}
	return false
}

func NewScoreService(logger *logrus.Logger, store db.Store) ScoreService {
	return &scoreService{
		logger,
		store,
	}
}
