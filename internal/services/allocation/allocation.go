package allocation

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/sirupsen/logrus"
)

type AllocationService interface {
	Create(string) (db.Allocation, error)
	GetAll() ([]db.Allocation, error)
	Delete(int32) error
}

type allocationService struct {
	logger *logrus.Logger
	store  db.Store
}

// Delete implements AllocationService.
func (i *allocationService) Delete(id int32) error {
  return i.store.DeleteAllocation(context.Background(), id)
}

// Create implements AllocationService.
func (i *allocationService) Create(name string) (db.Allocation, error) {
	return i.store.CreateAllocation(context.Background(), name)
}

// GetAll implements AllocationService.
func (i *allocationService) GetAll() ([]db.Allocation, error) {
	return i.store.GetAllocations(context.Background())
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
