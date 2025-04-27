package contributor

import (
	"context"
	"errors"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ContributorService interface {
	Create(db.CreateContributorParams) (uuid.UUID, error)
	GetByID(uuid.UUID) (db.ContributorAccountScore, error)
	GetAll() ([]db.ContributorAccountScore, error)
	GetScoreStatistics(uuid.UUID) (db.GetContributorScoreStatisticsRow, error)
	GetBestSellingScores(uuid.UUID) ([]db.GetContributorBestSellingScoresRow, error)
	Verify(uuid.UUID) error
	GetPaymentMethod(uuid.UUID) (db.PaymentMethod, error)
	UpsertPaymentMethod(db.UpsertContributorPaymentMethodParams) error
	GetPayments(uuid.UUID) ([]db.GetContributorPaymentsRow, error)
	GetPaymentStatistics(uuid.UUID) (db.GetContributorPaymentStatisticsRow, error)
}

type contributorService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetPaymentStatistics implements ContributorService.
func (c *contributorService) GetPaymentStatistics(id uuid.UUID) (db.GetContributorPaymentStatisticsRow, error) {
	ctx := context.Background()
	return c.store.GetContributorPaymentStatistics(ctx, id)
}

// GetPayments implements ContributorService.
func (c *contributorService) GetPayments(id uuid.UUID) ([]db.GetContributorPaymentsRow, error) {
	ctx := context.Background()
	return c.store.GetContributorPayments(ctx, id)
}

func (c *contributorService) checkContributorVerified(id uuid.UUID) (bool, error) {
	contributor, err := c.GetByID(id)
	if err != nil {
		return false, errors.New("Contributor not found")
	}

	if !contributor.IsVerified.Bool {
		return false, nil
	}

	return true, nil
}

// GetPaymentMethod implements ContributorService.
func (c *contributorService) GetPaymentMethod(id uuid.UUID) (db.PaymentMethod, error) {
	ok, err := c.checkContributorVerified(id)
	if !ok {
		return db.PaymentMethod{}, errors.New("Contributor is not verified")
	}
	if err != nil {
		return db.PaymentMethod{}, err
	}

	ctx := context.Background()
	return c.store.GetContributorPaymentMethod(ctx, id)
}

// CreatePaymentMethod implements ContributorService.
func (c *contributorService) UpsertPaymentMethod(params db.UpsertContributorPaymentMethodParams) error {
	ctx := context.Background()
	ok, err := c.checkContributorVerified(params.ContributorID)
	if !ok {
		return errors.New("Contributor is not verified")
	}
	if err != nil {
		return err
	}

	return c.store.UpsertContributorPaymentMethod(ctx, params)
}

// GetBestSellingScores implements ContributorService.
func (c *contributorService) GetBestSellingScores(contributorId uuid.UUID) ([]db.GetContributorBestSellingScoresRow, error) {
	ctx := context.Background()
	return c.store.GetContributorBestSellingScores(ctx, contributorId)
}

// GetStatistics implements ContributorService.-method
func (c *contributorService) GetScoreStatistics(contributorId uuid.UUID) (db.GetContributorScoreStatisticsRow, error) {
	ctx := context.Background()
	return c.store.GetContributorScoreStatistics(ctx, contributorId)
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
	contrib, _ := c.store.GetContributorById(ctx, params.ID)

	if contrib.Email == "" {
		c.store.CreateContributor(ctx, params)
	}

	c.store.UpdateAccountRole(ctx, db.UpdateAccountRoleParams{
		Rolename: db.RolenameContributor,
		ID:       params.ID,
	})

	return params.ID, nil
}

func NewContributorService(
	logger *logrus.Logger,
	store db.Store,
) ContributorService {
	return &contributorService{
		logger,
		store,
	}
}
