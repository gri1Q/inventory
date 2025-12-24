package model

import "time"

// Manufacturer — производитель товара.
// @desc Лучше держать отдельно, чем строкой в Item, чтобы избежать дублирования.
type Manufacturer struct {
	// @var string UUID
	ID int `json:"id"`
	// @var string Название производителя
	Name string `json:"name" `
	// @var time.Time CreatedAt
	CreatedAt time.Time `json:"created_at"`
	// @var time.Time UpdatedAt
	UpdatedAt time.Time `json:"updated_at"`
}
