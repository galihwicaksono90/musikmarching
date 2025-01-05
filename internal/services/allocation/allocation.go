package allocation

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AllocationService interface {
	Create(string) (db.Allocation, error)
	GetAll() ([]db.Allocation, error)
	Delete(int32) error
	CreateScoreAllocation(db.CreateScoreAllocationParams) error
	UpsertManyScoreAllocation(scoreId uuid.UUID, allocations []int32)
	DeleteScoreAllocationByScoreId(uuid.UUID) error
}

type allocationService struct {
	logger *logrus.Logger
	store  db.Store
}

// DeleteScoreAllocationByScoreId implements AllocationService.
func (s *allocationService) DeleteScoreAllocationByScoreId(scoreID uuid.UUID) error {
	return s.store.DeleteScoreAllocation(context.Background(), scoreID)
}

// UpsertManyScoreAllocation implements AllocationService.
func (s *allocationService) UpsertManyScoreAllocation(scoreId uuid.UUID, allocations []int32) {
	s.DeleteScoreAllocationByScoreId(scoreId)

	for _, a := range allocations {
		s.CreateScoreAllocation(db.CreateScoreAllocationParams{
			ScoreID:      scoreId,
			AllocationID: a,
		})
	}
}

// CreateScoreAllocation implements AllocationService.
func (a *allocationService) CreateScoreAllocation(params db.CreateScoreAllocationParams) error {
	return a.store.CreateScoreAllocation(context.Background(), params)
}

// Delete implements AllocationService.
func (a *allocationService) Delete(id int32) error {
	return a.store.DeleteAllocation(context.Background(), id)
}

// Create implements AllocationService.
func (a *allocationService) Create(name string) (db.Allocation, error) {
	return a.store.CreateAllocation(context.Background(), name)
}

// GetAll implements AllocationService.
func (a *allocationService) GetAll() ([]db.Allocation, error) {
	return a.store.GetAllocations(context.Background())
}

func NewAllocationService(
	logger *logrus.Logger,
	store db.Store,
) AllocationService {
	return &allocationService{
		logger,
		store,
	}
}
