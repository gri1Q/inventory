package model

import "time"

// User — учетная запись пользователя системы.
// @desc Хранит информацию для авторизации и привязку к роли.
type User struct {
	// @var string UUID PK
	ID string `json:"id"`
	// @var string Полное имя пользователя
	Name string `json:"name" `
	// @var string Email (используется для логина)
	Email string `json:"email" `
	// @var string Хэш пароля. НЕ возвращается в JSON.
	PasswordHash string `json:"-" `
	// @var int ID роли
	RoleID int `json:"role_id" `
	// @var bool Аккаунт активен
	Active bool `json:"active"`
	// @var time.Time CreatedAt
	CreatedAt time.Time `json:"created_at"`
	// @var time.Time UpdatedAt
	UpdatedAt time.Time `json:"updated_at"`
}
