package contributor

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ContributorService interface {
	SetUserAsContributor(id uuid.UUID) error
	GetUnverifiedContributor() ([]db.Contributor, error)
	VerifyContributor(id uuid.UUID) error
}

type contributorService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetUnverifiedContributor implements ContributorService.
func (c *contributorService) GetUnverifiedContributor() ([]db.Contributor, error) {
	return c.store.GetUnverifiedContributor(context.Background())
}

// VerifyContributor implements ContributorService.
func (c *contributorService) VerifyContributor(id uuid.UUID) error {
	return c.store.VerifyContributor(context.Background(), id)
}

// SetUserAsContributor implements ContributorService.
func (c *contributorService) SetUserAsContributor(id uuid.UUID) error {
	ctx := context.Background()

	_, err := c.store.CreateContributor(ctx, id)
	if err != nil {
		return err
	}

	_, err = c.store.UpdateAccountRole(ctx, db.UpdateAccountRoleParams{
		Rolename: db.RolenameContributor,
		ID:       id,
	})
	if err != nil {
		return err
	}

	return nil
}

func NewContributorService(logger *logrus.Logger, store db.Store) ContributorService {
	return &contributorService{
		logger,
		store,
	}
}
