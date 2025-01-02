package instrument

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/sirupsen/logrus"
)

type InstrumentService interface {
	Create(string) (db.Instrument, error)
	GetAll() ([]db.Instrument, error)
	Delete(int32) error
}

type instrumentService struct {
	logger *logrus.Logger
	store  db.Store
}

// Delete implements InstrumentService.
func (i *instrumentService) Delete(id int32) error {
  return i.store.DeleteInstrument(context.Background(), id)
}

// Create implements InstrumentService.
func (i *instrumentService) Create(name string) (db.Instrument, error) {
	return i.store.CreateInstrument(context.Background(), name)
}

// GetAll implements InstrumentService.
func (i *instrumentService) GetAll() ([]db.Instrument, error) {
	return i.store.GetInstruments(context.Background())
}

func NewInstrumentService(
	logger *logrus.Logger,
	store db.Store,
) InstrumentService {
	return &instrumentService{
		logger,
		store,
	}
}
