package service

import (
	"context"
	"inventory/internal/model"
)

// LocationRepository — контракт для работы с локациями.
type LocationRepository interface {
	// Create создает новую локацию.
	Create(ctx context.Context, l *model.Location) error
	// GetByID возвращает локацию по ID.
	GetByID(ctx context.Context, id int) (*model.Location, error)
	// Update обновляет данные локации.
	Update(ctx context.Context, l *model.Location) error
	// Delete удаляет локацию по ID.
	Delete(ctx context.Context, id int) error
	// GetAll возвращает список локаций с пагинацией.
	GetAll(ctx context.Context, limit, offset int) ([]*model.Location, error)
}
