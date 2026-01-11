package router

import (
	"inventory/internal/handler"
	"inventory/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(categoryHandler *handler.CategoryHandler) *gin.Engine {
	// Не используем gin.Default() чтобы контролировать middleware самостоятельно.
	r := gin.New()

	// Basic middlewares: logger + recovery. Можно заменить кастомным логгером.
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.MyGlobalCustomMiddleware())

	v1 := r.Group("/api/v1", middleware.MyGroupCustomMiddleware())
	{
		v1.GET("/categories", categoryHandler.ShowCategories)
		//TODO CSRF узнать
		v1.POST("/categories", middleware.MyRoutCustomMiddleware(), categoryHandler.CreateCategory)
	}

	return r
}
