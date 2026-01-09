package handler

import (
	"inventory/internal/dto"
	"net/http"

	"inventory/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}
func (cat *CategoryHandler) ShowCategories(c *gin.Context) {
	categories, err := cat.categoryService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (cat *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest

	// Валидация входных данных
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Вызов сервиса
	category, err := cat.categoryService.Create(c.Request.Context(), req)
	if err != nil {
		// Лучше логировать ошибку здесь
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create category",
		})
		return
	}

	c.JSON(http.StatusCreated, category)
}
