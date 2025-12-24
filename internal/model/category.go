package model

import "time"

// Category представляет категорию товара.
type Category struct {
	// @var string Уникальный идентификатор (UUID)
	ID int `json:"id"`
	// @var string Название категории
	Name string `json:"name"`
	// @var string Описание категории
	Description string `json:"description"`
	// @var *string ID родительской категории, nil если корневая
	ParentID *int `json:"parent_id"`
	// @var bool Активна ли категория
	Active bool `json:"active"`
	// @var time.Time Время создания
	CreatedAt time.Time `json:"created_at"`
	// @var time.Time Время последнего обновления
	UpdatedAt time.Time `json:"updated_at"`
}
