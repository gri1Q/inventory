package router

import (
	"inventory/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(categoryHandler *handler.CategoryHandler) *gin.Engine {
	// Не используем gin.Default() чтобы контролировать middleware самостоятельно.
	r := gin.New()

	// Basic middlewares: logger + recovery. Можно заменить кастомным логгером.
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		api.GET("/categories", categoryHandler.ShowCategories)
	}

	return r
}
