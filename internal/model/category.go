package model

import "time"

// Category описывает категорию товара.
type Category struct {
	// ID — первичный ключ (auto-increment).
	ID int `json:"id" `

	// Name — название категории.
	Name string `json:"name" `

	// Description — описание категории.
	Description string `json:"description,omitempty"`

	// Active — активна ли категория.
	Active bool `json:"active" `

	// CreatedAt — время создания записи.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt — время последнего обновления.
	UpdatedAt time.Time `json:"updated_at"`
}
