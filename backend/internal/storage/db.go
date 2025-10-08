// Package storage
//
// Defines DB operations that are available to use in handlers and services
package storage

import (
	"fmt"
	"time"

	"sso/internal/storage/models"
	"sso/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MustLoad(cfg *config.Config) *gorm.DB {
	dbConfig := gorm.Config{}
	if cfg.App.Production {
		dbConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DB.Host,
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Name,
			cfg.DB.Port), PreferSimpleProtocol: true,
	}), &dbConfig)
	if err != nil {
		panic(err)
	}

	d, err := db.DB()
	if err != nil {
		panic(err)
	}

	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(100)
	d.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&models.User{},
		&models.Session{},
		&models.Client{},
		&models.Permission{},
		&models.AuthRequest{},
	)

	return db
}
