package category

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CategoryService interface {
	Create(string) (db.Category, error)
	GetAll() ([]db.Category, error)
	Delete(int32) error
	CreateScoreCategory(db.CreateScoreCategoryParams) error
	UpsertManyScoreCategory(scoreId uuid.UUID, instruments []int32)
	DeleteScoreCategoryByScoreId(uuid.UUID) error
}

type categoryService struct {
	logger *logrus.Logger
	store  db.Store
}

// DeleteScoreCategoryByScoreId implements CategoryService.
func (c *categoryService) DeleteScoreCategoryByScoreId(id uuid.UUID) error {
	return c.store.DeleteScoreCategory(context.Background(), id)
}

// UpsertManyScoreCategory implements CategoryService.
func (c *categoryService) UpsertManyScoreCategory(scoreId uuid.UUID, categories []int32) {
	c.DeleteScoreCategoryByScoreId(scoreId)

	for _, x := range categories {
		c.CreateScoreCategory(db.CreateScoreCategoryParams{
			ScoreID:    scoreId,
			CategoryID: x,
		})
	}
}

// CreateScoreCategory implements CategoryService.
func (c *categoryService) CreateScoreCategory(params db.CreateScoreCategoryParams) error {
	return c.store.CreateScoreCategory(context.Background(), params)
}

// Delete implements CategoryService.
func (c *categoryService) Delete(id int32) error {
	return c.store.DeleteCategory(context.Background(), id)
}

// Create implements CategoryService.
func (c *categoryService) Create(name string) (db.Category, error) {
	return c.store.CreateCategory(context.Background(), name)
}

// GetAll implements CategoryService.
func (c *categoryService) GetAll() ([]db.Category, error) {
	return c.store.GetCategories(context.Background())
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
