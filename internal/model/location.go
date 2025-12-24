package model

import "time"

// Location обозначает место хранения/склад/офис/квартиру/кабинет.
// @desc Универсальная сущность для физических адресов и помещений.
type Location struct {
	// @var int UUID
	ID int `json:"id" `
	// @var string Название локации (например "Основной склад")
	Name string `json:"name" `
	// @var string Улица и номер дома
	Street string `json:"street,omitempty"`
	// @var string Номер кабинета/квартиры/ящика
	Room string `json:"room,omitempty" `
	// @var string Телефон контактного лица/склада
	Phone string `json:"phone,omitempty" `
	// @var string Email контакта
	Email string `json:"email,omitempty"`
	// @var bool Активна ли локация
	Active bool `json:"active"`
	// @var time.Time CreatedAt
	CreatedAt time.Time `json:"created_at"`
	// @var time.Time UpdatedAt
	UpdatedAt time.Time `json:"updated_at"`
}
