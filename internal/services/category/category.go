package category

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/sirupsen/logrus"
)

type CategoryService interface {
	Create(string) (db.Category, error)
	GetAll() ([]db.Category, error)
	Delete(int32) error
}

type categoryService struct {
	logger *logrus.Logger
	store  db.Store
}

// Delete implements CategoryService.
func (i *categoryService) Delete(id int32) error {
  return i.store.DeleteCategory(context.Background(), id)
}

// Create implements CategoryService.
func (i *categoryService) Create(name string) (db.Category, error) {
	return i.store.CreateCategory(context.Background(), name)
}

// GetAll implements CategoryService.
func (i *categoryService) GetAll() ([]db.Category, error) {
	return i.store.GetCategories(context.Background())
}

func NewCategoryService(
	logger *logrus.Logger,
	store db.Store,
) CategoryService {
	return &categoryService{
		logger,
		store,
	}
}
