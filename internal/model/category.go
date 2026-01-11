package model

import "time"

// Category описывает категорию товара.
type Category struct {
	// ID — первичный ключ (auto-increment).
	ID int

	// Name — название категории.
	Name string

	// Description — описание категории.
	Description string

	// Active — активна ли категория.
	Active bool `gorm:"default:true"`

	// CreatedAt — время создания записи.
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// UpdatedAt — время последнего обновления.
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
