package service

import (
	"context"
	"inventory/internal/model"
)

// UserRepository — контракт для работы с пользователями системы.
type UserRepository interface {
	// Create создает нового пользователя.
	Create(ctx context.Context, u *model.User) error
	// GetByID возвращает пользователя по ID.
	GetByID(ctx context.Context, id string) (*model.User, error)
	// GetByEmail возвращает пользователя по email.
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	// Update обновляет данные пользователя.
	Update(ctx context.Context, u *model.User) error
	// Delete удаляет пользователя по ID.
	Delete(ctx context.Context, id string) error
	// GetAll возвращает список пользователей с пагинацией.
	GetAll(ctx context.Context, limit, offset int) ([]*model.User, error)
}
