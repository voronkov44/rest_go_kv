package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"rest_go_kv/configs"
	"rest_go_kv/pkg/logger"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	logger.Info("Successfully connected to the database")
	return &Db{db}
}
