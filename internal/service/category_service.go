package service

import (
	"context"
	"inventory/internal/dto"
	"inventory/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(ctx context.Context, c *model.Category) error
	GetAll(ctx context.Context) ([]*model.Category, error)
	DB() *gorm.DB
}

type CategoryService struct {
	CategoryRepository CategoryRepository
}

func NewCategoryService(CategoryRepository CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: CategoryRepository}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]model.Category, error) {
	raw, err := s.CategoryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	categories := make([]model.Category, 0, len(raw))
	for _, r := range raw {
		categories = append(categories, *r)
	}
	return categories, nil
}

func (s *CategoryService) Create(ctx context.Context, req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {

	// начинаем транзакцию через GORM
	transaction := s.CategoryRepository.DB().Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}

	cat := &model.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	// создаем категорию в рамках транзакции
	if err := transaction.WithContext(ctx).Create(cat).Error; err != nil {
		transaction.Rollback() // откатываем при ошибке
		return nil, err
	}

	// формируем ответ
	res := &dto.CategoryResponse{
		ID:          cat.ID,
		Name:        cat.Name,
		Description: cat.Description,
		Active:      cat.Active,
		CreatedAt:   cat.CreatedAt,
		UpdatedAt:   cat.UpdatedAt,
	}

	if err := transaction.Commit().Error; err != nil {
		return nil, err
	}

	return res, nil
}
