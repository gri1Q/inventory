package router

import (
	"inventory/internal/config"
	"inventory/internal/handler"
	"inventory/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(categoryHandler *handler.CategoryHandler, cfg *config.Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.TimeoutMiddleware(cfg.Timeout))
	r.Use(middleware.MyGlobalCustomMiddleware())

	v1 := r.Group("/api/v1", middleware.MyGroupCustomMiddleware())
	{
		v1.GET("/categories", categoryHandler.ShowCategories)
		v1.POST("/categories", middleware.MyRoutCustomMiddleware(), categoryHandler.CreateCategory)
	}

	return r
}
