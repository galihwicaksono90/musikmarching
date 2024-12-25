package score

import (
	"context"
	"errors"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type ScoreService interface {
	GetAllPublic() ([]db.GetAllPublicScoresRow, error)
	GetManyByContributorId(account_id uuid.UUID) ([]db.Score, error)
	Create(model.CreateScoreDTO) (uuid.UUID, error)
	Update(uuid.UUID, model.UpdateScoreDTO) error
	GetManyVerified(db.GetVerifiedScoresParams) (*[]db.GetVerifiedScoresRow, error)
	GetVerifiedById(id uuid.UUID) (db.GetVerifiedScoreByIdRow, error)
	GetById(id uuid.UUID) (db.Score, error)
	GetManyByContirbutorID(db.GetScoresByContributorIDParams) ([]db.GetScoresByContributorIDRow, error)
	GetOneByContributorID(db.GetScoreByContributorIDParams) (db.GetScoreByContributorIDRow, error)
	GetAll() ([]db.Score, error)
	Verify(id uuid.UUID) error
}

type scoreService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetAllPublicScores implements ScoreService.
func (s *scoreService) GetAllPublic() ([]db.GetAllPublicScoresRow, error) {
	return s.store.GetAllPublicScores(context.Background(), db.GetAllPublicScoresParams{
		Pageoffset: 0,
		Pagelimit:  100,
	})
}

// GetOneByContributorID implements ScoreService.
func (s *scoreService) GetOneByContributorID(params db.GetScoreByContributorIDParams) (db.GetScoreByContributorIDRow, error) {
	return s.store.GetScoreByContributorID(context.Background(), params)
}

// GetByContirbutorID implements ScoreService.
func (s *scoreService) GetManyByContirbutorID(params db.GetScoresByContributorIDParams) ([]db.GetScoresByContributorIDRow, error) {
	return s.store.GetScoresByContributorID(context.Background(), params)
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
		Price: pgtype.Numeric{
			Int:   params.Price,
			Valid: true,
		},
		PdfUrl:        params.PdfUrl,
		PdfImageUrls:  params.PdfImageUrls,
		AudioUrl:      params.AudioUrl,
		ContributorID: params.ContributorID,
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
		PdfUrl:       params.PdfUrl,
		PdfImageUrls: params.PdfImageUrls,
		AudioUrl:     params.AudioUrl,
		ID:           scoreId,
	}); err != nil {
		return err
	}

	return nil
}

// GetScoresByContributorId implements ScoreService.
func (s *scoreService) GetScoresByContributorId(account_id uuid.UUID) ([]db.Score, error) {
	panic("unimplemented")
}

func NewScoreService(logger *logrus.Logger, store db.Store) ScoreService {
	return &scoreService{
		logger,
		store,
	}
}
