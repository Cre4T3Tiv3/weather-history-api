package main

import (
	"net/http"
	"os"
	"weather-history-api/configs"
	"weather-history-api/db"
	"weather-history-api/handlers"
	"weather-history-api/logger"
)

func loadConfig() (configs.Config, error) {
	configFile, err := os.Open("configs/config.json")
	defer configFile.Close()
	if err != nil {
		logger.Error.Println("Error opening configuration file:", err)
		return configs.Config{}, err
	}

	var config configs.Config
	config, err = configs.LoadConfig()
	if err != nil {
		logger.Error.Println("Error loading configuration:", err)
		return config, err
	}
	return config, nil
}

func main() {
	logger.InitLog()

	config, err := loadConfig()
	if err != nil {
		logger.Error.Fatal("Error loading configuration:", err)
	}

	if err := db.InitDB(config); err != nil {
		logger.Error.Fatal("Database connection error:", err)
	}
	db.MigrateSchema()

	r := handlers.SetupRoutes()

	http.Handle("/", r)
	logger.Info.Println("Starting server on port 8080...")
	logger.Error.Fatal(http.ListenAndServe(":8080", r))
}
