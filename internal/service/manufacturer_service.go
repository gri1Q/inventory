package service

import (
	"context"
	"inventory/internal/model"
)

// ManufacturerRepository — контракт для работы с производителями.
type ManufacturerRepository interface {
	// Create создает нового производителя.
	Create(ctx context.Context, m *model.Manufacturer) error
	// GetByID возвращает производителя по ID.
	GetByID(ctx context.Context, id int) (*model.Manufacturer, error)
	// Update обновляет данные производителя.
	Update(ctx context.Context, m *model.Manufacturer) error
	// Delete удаляет производителя по ID.
	Delete(ctx context.Context, id int) error
	// GetAll возвращает список производителей с пагинацией.
	GetAll(ctx context.Context, limit, offset int) ([]*model.Manufacturer, error)
}
