package score

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ScoreService interface {
	GetScoresByContributorId(account_id uuid.UUID) ([]db.Score, error)
	CreateScore() error
	GetVerifiedScores(db.GetVerifiedScoresParams) *[]db.GetVerifiedScoresRow
	GetVerifiedScoreById(id uuid.UUID) (db.GetVerifiedScoreByIdRow, error)
}

type scoreService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetVerifiedScoreById implements ScoreService.
func (s *scoreService) GetVerifiedScoreById(id uuid.UUID) (db.GetVerifiedScoreByIdRow, error) {
	return s.store.GetVerifiedScoreById(context.Background(), id)
}

// CreateScore implements ScoreService.
func (s *scoreService) CreateScore() error {
	panic("unimplemented")
}

// GetScoresByContributorId implements ScoreService.
func (s *scoreService) GetScoresByContributorId(account_id uuid.UUID) ([]db.Score, error) {
	panic("unimplemented")
}

// GetVerifiedScores implements ScoreService.
func (s *scoreService) GetVerifiedScores(params db.GetVerifiedScoresParams) *[]db.GetVerifiedScoresRow {
	scores, err := s.store.GetVerifiedScores(context.Background(), params)

	if err != nil {
		return &[]db.GetVerifiedScoresRow{}
	}
	return &scores
}

func NewScoreService(logger *logrus.Logger, store db.Store) ScoreService {
	return &scoreService{
		logger,
		store,
	}
}
