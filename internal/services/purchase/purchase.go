package purchase

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PurchaseService interface {
	PurchaseScore(accountID uuid.UUID, scoreID uuid.UUID) (uuid.UUID, error)
	GetPurchases(uuid.UUID) ([]db.Purchase, error)
}

type purchaseService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetPurchases implements PurchaseService.
func (p *purchaseService) GetPurchases(id uuid.UUID) ([]db.Purchase, error) {
	return p.store.GetPurchases(context.Background(), id)
}

// PurchaseScore implements PurchaseService.
func (p *purchaseService) PurchaseScore(accountID uuid.UUID, scoreID uuid.UUID) (uuid.UUID, error) {
	ctx := context.Background()
	var purchaseId uuid.UUID

	score, err := p.store.GetVerifiedScoreById(ctx, scoreID)
	if err != nil {
		return purchaseId, err
	}

	params := &db.CreatePurchaseParams{
		AccountID: accountID,
		Price:     score.Price,
		Title:     score.Title,
		ScoreID:   score.ID,
	}

	return p.store.CreatePurchase(context.Background(), *params)
}

func NewPurchaseService(logger *logrus.Logger, store db.Store) PurchaseService {
	return &purchaseService{
		logger,
		store,
	}
}
