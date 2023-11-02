package db

import (
	"fmt"
	"weather-history-api/configs"
	"weather-history-api/logger"
	"weather-history-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB(config configs.Config) error {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DBUser, config.DBPassword, config.DBName)

	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error.Println("Error opening database:", err)
		return err
	}

	logger.Info.Println("Database successfully connected!")
	return nil
}

func CloseDB() {
	sqlDB, err := DBConn.DB()
	if err != nil {
		logger.Error.Println("Error fetching underlying SQL database:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		logger.Error.Println("Error closing the database:", err)
	}
}

func MigrateSchema() {
	if err := DBConn.AutoMigrate(&models.Weather{}); err != nil {
		logger.Error.Println("Error migrating schema:", err)
	}
}
