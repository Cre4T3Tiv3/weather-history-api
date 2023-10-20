package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"weather-history-api/configs"
	"weather-history-api/db"
	"weather-history-api/handlers"
	"weather-history-api/models"
)

// loadConfig loads the configuration from a JSON file.
func loadConfig() (configs.Config, error) {
	var config configs.Config
	configFile, err := os.Open("configs/config.json")
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func main() {
	// Load configuration.
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize the database connection using the configuration values.
	db.InitDB(config)

	// Set up API routes.
	r := handlers.SetupRoutes()

	// Start the API server.
	http.Handle("/", r)
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func MigrateSchema() {
	err := db.DB.AutoMigrate(&models.Weather{})
	if err != nil {
		log.Fatal("Schema migration failed:", err)
	}
}
