package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"weather-history-api/configs"
	"weather-history-api/db"
	"weather-history-api/handlers"
	"weather-history-api/logger"
)

var localLogger *log.Logger

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
	logger.InitLogger()

	config, err := loadConfig()
	if err != nil {
		logger.Error.Fatal("Error loading configuration:", err)
	}

	db.InitDB(config)
	db.MigrateSchema()

	r := handlers.SetupRoutes()

	http.Handle("/", r)
<<<<<<< HEAD
	localLogger.Println("Starting server on port 8080...")
	logger.Info.Println("Starting server on port 8080...")
	logger.Error.Fatal(http.ListenAndServe(":8080", r))
=======
	log.Println("Starting server on port 8080...")
	logger.Println("Starting server on port 8080...")
	logger.Fatal(http.ListenAndServe(":8080", r))
}

func MigrateSchema() {
	err := db.DB.AutoMigrate(&models.Weather{})
	if err != nil {
		logger.Fatal("Schema migration failed:", err)
	}
>>>>>>> main
}
