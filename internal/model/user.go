package model

import "time"

// User — учётная запись пользователя системы.
type User struct {
	// ID — первичный ключ.
	ID string `json:"id"`
	// Name — полное имя пользователя.
	Name string `json:"name" `
	// Email — адрес электронной почты (используется для входа).
	Email string `json:"email" `
	// PasswordHash — хэш пароля; не возвращается в JSON.
	PasswordHash string `json:"-" `
	// RoleID — ссылка на роль (Role.ID).
	RoleID int `json:"role_id" `
	// Active — активен ли аккаунт.
	Active bool `json:"active"`
	// CreatedAt — время создания.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt — время обновления.
	UpdatedAt time.Time `json:"updated_at"`
}
