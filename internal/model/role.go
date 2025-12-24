package model

import "time"

// Role — роль пользователя (RBAC).
// @desc Простая таблица ролей. Для fine-grained правовая модель добавьте Permission и many2many.
type Role struct {
	// @var int
	ID int `json:"id"`
	// @var string Уникальное системное имя роли, например "admin"
	Name string `json:"name"`
	// @var string Описание роли
	Description string `json:"description,omitempty"`
	// @var time.Time CreatedAt
	CreatedAt time.Time `json:"created_at"`
	// @var time.Time UpdatedAt
	UpdatedAt time.Time `json:"updated_at"`
}
