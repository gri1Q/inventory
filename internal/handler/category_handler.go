package handler

import (
	"errors"
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
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	res, err := cat.categoryService.Create(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("already exists")):
			c.JSON(http.StatusConflict, gin.H{"error": "category already exists"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		}
		return
	}

	c.JSON(http.StatusCreated, res)
}
