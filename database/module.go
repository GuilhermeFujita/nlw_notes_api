package database

import (
	"context"
	"fmt"
	"time"

	"github.com/GuilhermeFujita/nlw_notes_api/config"
	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Provide(NewDatabase)

func NewDatabase(lc fx.Lifecycle, cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Opening DB: %w", err)
	}

	if err := db.AutoMigrate(&model.Note{}); err != nil {
		return nil, fmt.Errorf("Auto migrate error %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("SQL DB error %w", err)
	}

	// Pool Configurations
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return sqlDB.Ping()
		},
		OnStop: func(ctx context.Context) error {
			return sqlDB.Close()
		},
	})

	return db, nil
}
