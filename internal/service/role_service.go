package service

import (
	"context"
	"inventory/internal/model"
)

// RoleRepository — контракт для работы с ролями пользователей.
type RoleRepository interface {
	// Create создает новую роль.
	Create(ctx context.Context, r *model.Role) error
	// GetByID возвращает роль по ID.
	GetByID(ctx context.Context, id int) (*model.Role, error)
	// Update обновляет данные роли.
	Update(ctx context.Context, r *model.Role) error
	// Delete удаляет роль по ID.
	Delete(ctx context.Context, id int) error
	// GetAll возвращает список ролей с пагинацией.
	GetAll(ctx context.Context, limit, offset int) ([]*model.Role, error)
}
