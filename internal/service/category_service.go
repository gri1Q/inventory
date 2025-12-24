package service

import (
	"context"
	"inventory/internal/model"
)

// CategoryRepository — контракт для работы с категориями.
type CategoryRepository interface {
	// Create создает новую категорию.
	Create(ctx context.Context, c *model.Category) error
	// GetByID возвращает категорию по ID.
	GetByID(ctx context.Context, id int) (*model.Category, error)
	// Update обновляет данные категории.
	Update(ctx context.Context, c *model.Category) error
	// Delete удаляет категорию по ID.
	Delete(ctx context.Context, id int) error
	// GetAll возвращает список категорий с пагинацией.
	GetAll(ctx context.Context, limit, offset int) ([]*model.Category, error)
}

type CategoryService struct {
	CategoryRepository CategoryRepository
}

func (cs CategoryService) CreateCategory(ctx context.Context, category model.Category) (model.Category, error) {

}
