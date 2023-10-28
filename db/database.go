package db

import (
	"fmt"
	"log"

	"weather-history-api/configs"
	"weather-history-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// SetDBConnection sets the global DB connection with the provided database instance.
func SetDBConnection(database *gorm.DB) {
	DBConn = database
}

// GetDBConnection retrieves the global DB connection.
func GetDBConnection() *gorm.DB {
	return DBConn
}

func InitDB(config configs.Config) error {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DBUser, config.DBPassword, config.DBName)

	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	log.Println("Database successfully connected!")
	return nil
}

func CloseDB() {
	sqlDB, err := DBConn.DB()
	if err != nil {
		log.Println("Error fetching underlying SQL database:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Error closing the database:", err)
	}
}

func MigrateSchema() error {
	if err := DBConn.AutoMigrate(&models.Weather{}); err != nil {
		return fmt.Errorf("error migrating schema: %v", err)
	}

	return nil
}
