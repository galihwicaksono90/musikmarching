package profile

import (
	"context"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ProfileService interface {
	GetProfileByAccountId(account_id uuid.UUID) (*model.Profile, error)
}

type profileService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetScoresByAccountId implements ProfileService.
func (p *profileService) GetProfileByAccountId(account_id uuid.UUID) (*model.Profile, error) {

	profile, err := p.store.GetProfileByAccountId(context.Background(), account_id)
	if err != nil {
		return nil, err
	}

	uploadedScores, err := p.store.GetScoresByProfile(context.Background(), account_id)
	if err != nil {
		p.logger.Errorln(err)
	}

	return &model.Profile{
		Profile:        &profile,
		UploadedScores: &uploadedScores,
	}, nil
}

func NewProfileService(logger *logrus.Logger, store db.Store) ProfileService {
	return &profileService{
		logger,
		store,
	}
}
