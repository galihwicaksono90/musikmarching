package instrument

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type InstrumentService interface {
	Create(string) (db.Instrument, error)
	GetAll() ([]db.Instrument, error)
	Delete(int32) error
	CreateScoreInstrument(db.CreateScoreInstrumentParams) error
	UpsertManyScoreInstrument(scoreId uuid.UUID, instruments []int32)
	DeleteScoreInstrumentByScoreId(uuid.UUID) error
}

type instrumentService struct {
	logger *logrus.Logger
	store  db.Store
}

func (i *instrumentService) UpsertManyScoreInstrument(scoreId uuid.UUID, instruments []int32) {
	i.DeleteScoreInstrumentByScoreId(scoreId)

	for _, c := range instruments {
		i.CreateScoreInstrument(db.CreateScoreInstrumentParams{
			ScoreID:      scoreId,
			InstrumentID: c,
		})
	}
}

// DeletScoreInstrument implements InstrumentService.
func (i *instrumentService) DeleteScoreInstrumentByScoreId(scoreId uuid.UUID) error {
	return i.store.DeleteScoreInstrument(context.Background(), scoreId)
}

// CreateScoreInstrument implements InstrumentService.
func (i *instrumentService) CreateScoreInstrument(params db.CreateScoreInstrumentParams) error {
	return i.store.CreateScoreInstrument(context.Background(), params)
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
