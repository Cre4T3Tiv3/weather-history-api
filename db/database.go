package db

import (
	"log"
	"os"
	"weather-history-api/configs"
	"weather-history-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var logger *log.Logger

func init() {
	// Initialize the logger.
	logFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening or creating log file: %v", err)
	}
	// Create a logger instance to write logs to the file.
	logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
}

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
	logger.Println("Database successfully connected!")
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
