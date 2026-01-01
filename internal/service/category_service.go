package service

import (
	"context"
	"inventory/internal/model"
	"time"
)

// CategoryRepository — контракт для работы с категориями.
//type CategoryRepository interface {
// Create создает новую категорию.
//Create(ctx context.Context, c *model.Category) error
//// GetByID возвращает категорию по ID.
//GetByID(ctx context.Context, id int) (*model.Category, error)
//// Update обновляет данные категории.
//Update(ctx context.Context, c *model.Category) error
//// Delete удаляет категорию по ID.
//Delete(ctx context.Context, id int) error
//// GetAll возвращает список категорий с пагинацией.
//GetAll(ctx context.Context, limit, offset int) ([]*model.Category, error)
//}

type CategoryService struct {
	//CategoryRepository CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]model.Category, error) {
	now := time.Now()

	categories := []model.Category{
		{
			ID:          1,
			Name:        "Электроника",
			Description: "Телефоны, ноутбуки и гаджеты",
			ParentID:    nil,
			Active:      true,
			CreatedAt:   now.Add(-72 * time.Hour),
			UpdatedAt:   now,
		},
		{
			ID:          2,
			Name:        "Одежда",
			Description: "Мужская и женская одежда",
			ParentID:    nil,
			Active:      true,
			CreatedAt:   now.Add(-48 * time.Hour),
			UpdatedAt:   now,
		},
		{
			ID:          3,
			Name:        "Смартфоны",
			Description: "Мобильные телефоны и аксессуары",
			ParentID:    intPtr(1), // дочерняя категория от "Электроника"
			Active:      true,
			CreatedAt:   now.Add(-24 * time.Hour),
			UpdatedAt:   now,
		},
	}

	return categories, nil
}
func intPtr(v int) *int {
	return &v
}
