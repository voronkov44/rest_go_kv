package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"rest_go_kv/internal/orders"
	"rest_go_kv/internal/users"
	"rest_go_kv/pkg/logger"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file: %v", err)
		os.Exit(1)
	}
	logger.Info("Environment variables loaded successfully")

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	logger.Info("Successfully connected to the database")

	err = db.AutoMigrate(&users.User{}, &orders.Order{})
	if err != nil {
		logger.Error("Error migrating database: %v", err)
		os.Exit(1)
	}
	logger.Info("Database migration completed successfully")

}
