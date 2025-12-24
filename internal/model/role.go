package model

import "time"

// Role — роль пользователя
type Role struct {
	// ID — первичный ключ.
	ID int `json:"id"`
	// Name — системное имя роли (например "admin", "manager").
	Name string `json:"name"`
	// Description — описание роли.
	Description string `json:"description,omitempty"`
	// CreatedAt — время создания.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt — время обновления.
	UpdatedAt time.Time `json:"updated_at"`
}
