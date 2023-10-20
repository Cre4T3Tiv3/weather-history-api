package db

import (
	"log"
	"weather-history-api/configs"
	"weather-history-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection using the provided configuration.
func InitDB(config configs.Config) {
	// Construct the database connection string (DSN).
	dsn := "user=" + config.DBUser + " password=" + config.DBPassword + " dbname=" + config.DBName + " sslmode=disable"

	var err error
	// Open a connection to the database.
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database successfully connected!")

	// Auto migration for Weather model.
	DB.AutoMigrate(&models.Weather{})
}

// CloseDB closes the underlying SQL database connection.
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Error fetching underlying SQL database:", err)
		return
	}
	sqlDB.Close()
}
