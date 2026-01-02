package postgredb

import (
	"fmt"
	"inventory/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect открывает соединение gorm.DB по данным из cfg.DatabaseServer.
func Connect(cfg *config.Config) (*gorm.DB, error) {
	// если хотите, можно поддержать DB_URL: если задан - используем его
	// dbURL := os.Getenv("DB_URL")

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		cfg.DatabaseServer.Host,
		cfg.DatabaseServer.Port,
		cfg.DatabaseServer.User,
		cfg.DatabaseServer.Password,
		cfg.DatabaseServer.DBName,
		cfg.DatabaseServer.SSLMode,
	)

	gormCfg := &gorm.Config{
		// тут можно настроить logger/опции
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		return nil, err
	}

	// получить базовый *sql.DB чтобы настроить пул
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Ping чтобы убедиться, что соединение живо
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
