package payment

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PaymentService interface {
	Create(uuid.UUID) error
}

type paymentService struct {
	loggger *logrus.Logger
	store   db.Store
}

// Create implements PaymentService.
func (p *paymentService) Create(id uuid.UUID) error {
	ctx := context.Background()
	return p.store.CreatePayment(ctx, id)
}

func NewPaymentService(
	logger *logrus.Logger,
	store db.Store,
) PaymentService {
	return &paymentService{logger, store}
}
