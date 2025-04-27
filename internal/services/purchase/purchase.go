package purchase

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PurchaseService interface {
	PurchaseScore(uuid.UUID, uuid.UUID) (uuid.UUID, error)
	GetAll() ([]db.GetAllPurchasesRow, error)
	Verify(purchase_id uuid.UUID) error
	GetPurchasesByAccountID(uuid.UUID) ([]db.Purchase, error)
	GetPurchaseByID(db.GetPurchaseByIdParams) (db.Purchase, error)
	GetPurchasedScoreById(uuid.UUID, uuid.UUID) (db.GetPurchasedScoreByIdRow, error)
	UpdatePurchaseProof(db.UpdatePurchaseProofParams) error
}

type purchaseService struct {
	logger *logrus.Logger
	store  db.Store
}

// GetPurchasedScoreById implements PurchaseService.
func (p *purchaseService) GetPurchasedScoreById(purchaseID uuid.UUID, accountID uuid.UUID) (db.GetPurchasedScoreByIdRow, error) {
	ctx := context.Background()
	purchase, err := p.GetPurchaseByID(db.GetPurchaseByIdParams{
		ID:        purchaseID,
		AccountID: accountID,
	})
	if err != nil || !purchase.IsVerified {
		return db.GetPurchasedScoreByIdRow{}, err
	}
	return p.store.GetPurchasedScoreById(ctx, purchase.ScoreID)
}

// Verify implements PurchaseService.
func (p *purchaseService) Verify(id uuid.UUID) error {
	ctx := context.Background()
	return p.store.ExecTx(ctx, func(q *db.Queries) error {
		_, err := q.VerifyPurchase(ctx, id)

		if err != nil {
			p.logger.Errorln(err)
			return err
		}

		if err := p.store.CreatePayment(ctx, id); err != nil {
			return err
		}
		return nil
	})
}

// GetAll implements PurchaseService.
func (p *purchaseService) GetAll() ([]db.GetAllPurchasesRow, error) {
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
