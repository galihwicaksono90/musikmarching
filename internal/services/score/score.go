package score

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Profile struct {
	ID             uuid.UUID  `json:"id"`
	AccountID      uuid.UUID  `json:"account_id"`
	UploadedScores []db.Score `json:"uploaded_scores"`
}

type ScoreService interface {
	GetScoresByContributorId(account_id uuid.UUID) ([]db.Score, error)
	CreateScore(db.CreateScoreParams) error
	GetVerifiedScores() (*[]db.GetVerifiedScoresRow, error)
}

type scoreService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetVerifiedScores implements ScoreService.
func (p *scoreService) GetVerifiedScores() (*[]db.GetVerifiedScoresRow, error) {
	scores, err := p.store.GetVerifiedScores(context.Background())
	if err != nil {
		return &[]db.GetVerifiedScoresRow{}, err
	}
	return &scores, nil
}

// CreateScore implements ScoreService.
func (p *scoreService) CreateScore(params db.CreateScoreParams) error {
	return p.store.CreateScore(context.Background(), params)
}

// GetScoresByAccountId implements ProfileService.
func (p *scoreService) GetScoresByContributorId(account_id uuid.UUID) ([]db.Score, error) {
	return p.store.GetScoresByContributorId(context.Background(), account_id)
}

func NewScoreService(logger *logrus.Logger, store db.Store) ScoreService {
	return &scoreService{
		logger,
		store,
	}
}
