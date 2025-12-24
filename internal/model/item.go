package model

import "time"

// Item — карточка товара (SKU).
type Item struct {
	// ID — первичный ключ.
	ID int `json:"id"`

	// SKU — уникальный артикул/код товара.
	SKU string `json:"sku"`

	// Name — наименование товара.
	Name string `json:"name"`

	// Description — описание товара.
	Description string `json:"description,omitempty"`

	// CategoryID — ссылка на категорию Category.ID.
	CategoryID int `json:"category_id,omitempty"`

	// UoMID — ссылка на единицу измерения.
	UoMID int `json:"uom_id,omitempty"`

	// ManufacturerID — ссылка на производителя Manufacturer.ID.
	ManufacturerID int `json:"manufacturer_id,omitempty"`

	// Model — модель/номер модели производителя.
	Model string `json:"model,omitempty"`

	// IsSerialTracked — отслеживается ли по серийному номеру.
	IsSerialTracked bool `json:"is_serial_tracked"`

	// IsBatchTracked — отслеживается ли по партии/лоту.
	IsBatchTracked bool `json:"is_batch_tracked"`

	// DefaultCost — стандартная стоимость в минимальных единицах (копейки/центы).
	DefaultCost int64 `json:"default_cost"`

	// Active — активен ли товар.
	Active bool `json:"active"`

	// CreatedAt — время создания.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt — время обновления.
	UpdatedAt time.Time `json:"updated_at"`
}
