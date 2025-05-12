package contributorapply

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ContributorApplyService interface {
	Apply(db.CreateContributorApplyParams) (db.ContributorApply, error)
	GetByAccountID(uuid.UUID) (db.ContributorApply, error)
	Update(db.UpdateContributorApplyParams) error
	GetAll() ([]db.ContributorApply, error)
}

type contributorApplyService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetAll implements ContributorApplyService.
func (s *contributorApplyService) GetAll() ([]db.ContributorApply, error) {
	return s.store.GetContributorApplications(context.Background())
}

// Update implements ContributorApplyService.
func (s *contributorApplyService) Update(params db.UpdateContributorApplyParams) error {
	ctx := context.Background()
	return s.store.UpdateContributorApply(ctx, params)
}

// GetContributorApplyByAccountID implements ContributorApplyService.
func (s *contributorApplyService) GetByAccountID(accountID uuid.UUID) (db.ContributorApply, error) {
	ctx := context.Background()
	return s.store.GetContributorApplyByAccountID(ctx, accountID)
}

// ApplyContributor implements AccountService.
func (s *contributorApplyService) Apply(params db.CreateContributorApplyParams) (db.ContributorApply, error) {
	ctx := context.Background()
	return s.store.CreateContributorApply(ctx, params)
}

func NewContributorApplyService(logger *logrus.Logger, store db.Store) ContributorApplyService {
	return &contributorApplyService{
		logger,
		store,
	}
}
