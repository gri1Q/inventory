package model

import "time"

// Item — карточка товара.
// @desc Содержит информацию о товаре, производителе, учёте серий/партии и т.д.

type Item struct {
	// @var string UUID
	ID string `json:"id"`
	// @var string Код товара (SKU), уникальный
	SKU string `json:"sku"`
	// @var string Название товара
	Name string `json:"name"`
	// @var string Описание
	Description string `json:"description"`
	// @var string Ссылка на категорию (Category.ID)
	CategoryID int `json:"category_id"`
	// @var int ID единицы измерения (UoM)
	UoMID int `json:"uom_id"`
	// @var string Ссылка на производителя (Manufacturer.ID)
	ManufacturerID string `json:"manufacturer_id"`
	// @var string Модель/номер модели производителя
	Model string `json:"model"`
	// @var int Идентификатор нахождения
	LocationID int `json:"location_id"`
	// @var bool Отслеживается ли по серийному номеру
	IsSerialTracked bool `json:"is_serial_tracked"`
	// @var bool Отслеживается ли по партии/лоту
	IsBatchTracked bool `json:"is_batch_tracked"`
	// @var int64 Стоимость в минимальных единицах (копейках/центах). НЕ использовать float.
	DefaultCost int64 `json:"default_cost"`
	// @var bool Активность товара
	Active bool `json:"active"`
	// @var time.Time CreatedAt
	CreatedAt time.Time `json:"created_at"`
	// @var time.Time UpdatedAt
	UpdatedAt time.Time `json:"updated_at"`
}
