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
	GetScoresByAccountId(account_id uuid.UUID) ([]db.Score, error)
	CreateScore(db.CreateScoreParams) error
}

type scoreService struct {
	logger *logrus.Logger
	store  db.Store
}

// CreateScore implements ScoreService.
func (p *scoreService) CreateScore(params db.CreateScoreParams) error {
	return p.store.CreateScore(context.Background(), params)
}

// GetScoresByAccountId implements ProfileService.
func (p *scoreService) GetScoresByAccountId(account_id uuid.UUID) ([]db.Score, error) {
	return p.store.GetScoresByProfile(context.Background(), account_id)
}

func NewProfileService(logger *logrus.Logger, store db.Store) ScoreService {
	return &scoreService{
		logger,
		store,
	}
}
