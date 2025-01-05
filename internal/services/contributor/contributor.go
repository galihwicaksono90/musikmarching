package contributor

import (
	"context"
	"galihwicaksono90/musikmarching-be/internal/services/instrument"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ContributorService interface {
	Create(db.CreateContributorParams) (uuid.UUID, error)
	GetByID(uuid.UUID) (db.ContributorAccountScore, error)
	GetAll() ([]db.ContributorAccountScore, error)
	Verify(uuid.UUID) error
}

type contributorService struct {
	logger *logrus.Logger
	store  db.Store
	instrument instrument.InstrumentService
}

// Verify implements ContributorService.
func (c *contributorService) Verify(id uuid.UUID) error {
	ctx := context.Background()
	return c.store.VerifyContributor(ctx, id)
}

// GetAll implements ContributorService.
func (c *contributorService) GetAll() ([]db.ContributorAccountScore, error) {
	ctx := context.Background()
	return c.store.GetAllContributors(ctx)
}

// GetByID implements ContributorService.
func (c *contributorService) GetByID(id uuid.UUID) (db.ContributorAccountScore, error) {
	ctx := context.Background()
	return c.store.GetContributorById(ctx, id)
}

// Create implements ContributorService.
func (c *contributorService) Create(params db.CreateContributorParams) (uuid.UUID, error) {
	ctx := context.Background()

	_, err := c.store.CreateContributor(ctx, params)

	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = c.store.UpdateAccountRole(ctx, db.UpdateAccountRoleParams{
		Rolename: db.RolenameContributor,
		ID:       params.ID,
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return params.ID, nil
}

func NewContributorService(
	logger *logrus.Logger,
	store db.Store,
	instrument instrument.InstrumentService,
) ContributorService {
	return &contributorService{
		logger,
		store,
		instrument,
	}
}
