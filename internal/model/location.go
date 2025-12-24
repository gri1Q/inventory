package model

import "time"

// Location — место хранения: склад, офис, квартира, кабинет и т.п.
type Location struct {
	// ID — первичный ключ.
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`

	// Name — удобное название (например "Основной склад").
	Name string `json:"name"`

	// Street — улица и номер дома.
	Street string `json:"street,omitempty"`

	// Building — корпус/строение (опционально).
	Building string `json:"building,omitempty"`

	// Floor — этаж (опционально).
	Floor string `json:"floor,omitempty"`

	// Room — номер кабинета/квартиры/ячейки (опционально).
	Room string `json:"room,omitempty"`

	// Phone — контактный телефон.
	Phone string `json:"phone,omitempty"`

	// Email — контактный email.
	Email string `json:"email,omitempty"`

	// Active — активна ли локация.
	Active bool `json:"active"`

	// CreatedAt — время создания.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt — время обновления.
	UpdatedAt time.Time `json:"updated_at"`
}
