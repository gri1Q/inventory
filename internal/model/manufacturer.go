package model

import "time"

// Manufacturer — производитель товара.
type Manufacturer struct {
	// ID — первичный ключ.
	ID int `json:"id"`

	// Name — название производителя.
	Name string `json:"name"`

	// CreatedAt — время создания.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt — время обновления.
	UpdatedAt time.Time `json:"updated_at"`
}
