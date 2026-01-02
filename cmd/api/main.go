package main

import (
	"fmt"
	"inventory/internal/config"
	"inventory/internal/handler"
	"inventory/internal/router"
	"inventory/internal/service"
	"inventory/pkg/client/postgredb"
	"log"
)

func main() {
	// Подключение Config
	cfg := config.MustLoad()

	// Подключаемся к БД
	db, err := postgredb.Connect(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Гарантируем корректное закрытие соединения при завершении
	sqlDB, errorsDB := db.DB()
	if errorsDB != nil {
		log.Fatal(errorsDB)
	}
	defer sqlDB.Close()

	fmt.Println("✅ Подключение к PostgreSQL успешно установлено")

	// Сборка роутера с внедрёнными сервисами
	categoryService := service.NewCategoryService()
	categoryHandler := handler.NewCategoryHandler(categoryService)
	router := router.SetupRouter(categoryHandler)
	router.Run(cfg.HTTPServer.Addr)

	//TODO: init logger: slog
	//TODO init storage
}
