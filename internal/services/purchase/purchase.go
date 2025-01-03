package purchase

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PurchaseService interface {
	PurchaseScore(uuid.UUID, uuid.UUID) (uuid.UUID, error)
	GetAll() ([]db.Purchase, error)
	Verify(uuid.UUID) error
	GetPurchasesByAccountID(uuid.UUID) ([]db.Purchase, error)
	GetPurchaseByID(db.GetPurchaseByIdParams) (db.Purchase, error)
	UpdatePurchaseProof(db.UpdatePurchaseProofParams) error
}

type purchaseService struct {
	logger *logrus.Logger
	store  db.Store
}

// Verify implements PurchaseService.
func (p *purchaseService) Verify(id uuid.UUID) error {
	ctx := context.Background()
	return p.store.VerifyPurchase(ctx, id)
}

// GetAll implements PurchaseService.
func (p *purchaseService) GetAll() ([]db.Purchase, error) {
	ctx := context.Background()
	return p.store.GetAllPurchases(ctx)
}

// UpdatePurchaseProof implements PurchaseService.
func (p *purchaseService) UpdatePurchaseProof(params db.UpdatePurchaseProofParams) error {
	ctx := context.Background()

	err := p.store.UpdatePurchaseProof(ctx, params)
	return err
}

// GetPurchaseByID implements PurchaseService.
func (p *purchaseService) GetPurchaseByID(params db.GetPurchaseByIdParams) (db.Purchase, error) {
	ctx := context.Background()
	return p.store.GetPurchaseById(ctx, params)
}

// GetPurchases implements PurchaseService.
func (p *purchaseService) GetPurchasesByAccountID(id uuid.UUID) ([]db.Purchase, error) {
	ctx := context.Background()
	res, err := p.store.GetPurchasesByAccountId(ctx, id)
	if err != nil {
		p.logger.Errorln(err)
		return []db.Purchase{}, err
	}

	p.logger.Println(res)
	return res, nil
}

// PurchaseScore implements PurchaseService.
func (p *purchaseService) PurchaseScore(userID uuid.UUID, scoreID uuid.UUID) (uuid.UUID, error) {
	ctx := context.Background()

	score, err := p.store.GetVerifiedScoreById(ctx, scoreID)
	if err != nil {
		return uuid.New(), err
	}

	params := db.CreatePurchaseParams{
		AccountID: userID,
		Price:     score.Price,
		Title:     score.Title,
		ScoreID:   scoreID,
	}

	purchaseID, err := p.store.CreatePurchase(ctx, params)
	if err != nil {
		return uuid.New(), err
	}

	return purchaseID, nil
}

func NewPurchaseService(logger *logrus.Logger, store db.Store) PurchaseService {
	return &purchaseService{
		logger,
		store,
	}
}
