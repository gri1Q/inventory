package repository

import (
	"context"
	"inventory/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, category *model.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]*model.Category, error) {
	var categories []*model.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}
