package db

import (
	"fmt"
	"log"
<<<<<<< HEAD

=======
	"os"
>>>>>>> main
	"weather-history-api/configs"
	"weather-history-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

<<<<<<< HEAD
var DBConn *gorm.DB

// SetDBConnection sets the global DB connection with the provided database instance.
func SetDBConnection(database *gorm.DB) {
	DBConn = database
=======
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
>>>>>>> main
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
<<<<<<< HEAD

	log.Println("Database successfully connected!")
	return nil
=======
	log.Println("Database successfully connected!")
	logger.Println("Database successfully connected!")
	// Auto migration for Weather model.
	DB.AutoMigrate(&models.Weather{})
>>>>>>> main
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
