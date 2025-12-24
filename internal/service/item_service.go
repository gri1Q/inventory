package service

import (
	"context"
	"inventory/internal/model"
)

// ItemInterface - контракт для работы с товарами.
// Все методы возвращают ошибку (error) - это идиоматичный подход в Go.
type ItemRepository interface {
	// Create создает новый товар.
	Create(ctx context.Context, item *model.Item) error
	// GetByID возвращает товар по ID.
	GetByID(ctx context.Context, id int) (*model.Item, error)
	// Delete удаляет товар по ID.
	Delete(ctx context.Context, id int) error
	// GetAll возвращает список товаров с пагинацией и фильтрацией.
	GetAll(ctx context.Context, limit, offset int, filters map[string]interface{})
}
