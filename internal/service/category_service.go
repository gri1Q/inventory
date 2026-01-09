package service

import (
	"context"
	"inventory/internal/dto"
	"inventory/internal/model"
	"time"
)

type CategoryRepository interface {
	Create(ctx context.Context, c *model.Category) error
	GetAll(ctx context.Context) ([]*model.Category, error)
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
	now := time.Now()
	active := true
	if req.Active != nil {
		active = *req.Active
	}

	cat := &model.Category{
		Name:        req.Name,
		Description: req.Description,
		Active:      active,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.CategoryRepository.Create(ctx, cat); err != nil {
		return nil, err
	}

	res := &dto.CategoryResponse{
		ID:          cat.ID,
		Name:        cat.Name,
		Description: cat.Description,
		Active:      cat.Active,
		CreatedAt:   cat.CreatedAt,
		UpdatedAt:   cat.UpdatedAt,
	}

	return res, nil
}
