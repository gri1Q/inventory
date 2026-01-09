package dto

import "time"

// Входящий запрос от клиента
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Description string `json:"description"`
	Active      *bool  `json:"active"` // pointer, чтобы отличать "отсутствует" от false
}

// Выходящий ответ клиенту
type CategoryResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
