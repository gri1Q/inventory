package main

import (
	"inventory/internal/config"
	"inventory/internal/handler"
	"inventory/internal/router"
	"inventory/internal/service"
)

func main() {
	cfg := config.MustLoad()

	// Сборка роутера с внедрёнными сервисами
	categoryService := service.NewCategoryService()
	categoryHandler := handler.NewCategoryHandler(categoryService)
	router := router.SetupRouter(categoryHandler)
	router.Run(cfg.HTTPServer.Addr)

	//TODO: init logger: slog
	//TODO init storage
}
