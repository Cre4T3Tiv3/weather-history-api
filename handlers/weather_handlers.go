package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"weather-history-api/db"
	"weather-history-api/models"
	"weather-history-api/utils"

	"github.com/gorilla/mux"
)

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

// CreateWeather handles the creation of weather records via HTTP POST request.
func CreateWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST request to /api/weather.")
	logger.Println("Received POST request to /api/weather.")

	// Parse the request body.
	var weather models.Weather
	err := json.NewDecoder(r.Body).Decode(&weather)
	if err != nil {
		log.Println("Error decoding request body:", err)
		logger.Println("Error decoding request body:", err)
		utils.HttpError(w, "Error decoding request body", err, http.StatusBadRequest)
		return
	}

	// Convert temperature from Celsius to Fahrenheit.
	weather.TemperatureF = utils.CelsiusToFahrenheit(weather.TemperatureF)

	log.Println("Parsed weather data:", weather)
	logger.Println("Parsed weather data:", weather)

	// Validate the weather data (you can add more validation as needed).
	if weather.Date.IsZero() || weather.Location == "" || weather.TemperatureF == 0 {
		log.Println("Invalid weather data.")
		logger.Println("Invalid weather data.")
		utils.HttpError(w, "Invalid weather data", err, http.StatusBadRequest)
		return
	}

	// Insert the new weather record into the database.
	if err := db.DB.Create(&weather).Error; err != nil {
		if strings.Contains(err.Error(), "idx_weathers_location") {
			log.Println("Duplicate entry for location:", weather.Location)
			logger.Println("Duplicate entry for location:", weather.Location)
			utils.HttpError(w, "Weather data already exists for this location.", err, http.StatusBadRequest)
		} else {
			log.Println("Error inserting into database:", err)
			logger.Println("Error inserting into database:", err)
			utils.HttpError(w, "Error inserting into database:", err, http.StatusInternalServerError)
		}
		return
	}

	log.Println("Successfully inserted weather data into the database.")
	logger.Println("Successfully inserted weather data into the database.")

	// Respond with the newly created weather record.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(weather) // Note the = instead of :=
	if err != nil {
		log.Println("Error encoding response:", err)
		logger.Println("Error encoding response:", err)
		utils.HttpError(w, "Error encoding response", err, http.StatusInternalServerError)
		return
	}
	jsonResponse(w, weather, http.StatusCreated)
	log.Println("Response sent successfully.")
	logger.Println("Response sent successfully.")
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request for all weather records.")
	logger.Println("Received GET request for all weather records.")

	var weathers []models.Weather
	if err := db.DB.Find(&weathers).Error; err != nil {
		httpError(w, "Error fetching weather records", err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, weathers, http.StatusOK)
}

func GetWeatherByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request for a specific weather record.")
	logger.Println("Received GET request for a specific weather record.")

	params := mux.Vars(r)
	id := params["id"]

	var weather models.Weather
	if err := db.DB.First(&weather, id).Error; err != nil {
		httpError(w, "Record not found", err, http.StatusNotFound)
		return
	}

	jsonResponse(w, weather, http.StatusOK)
}

func UpdateWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received PUT request to update weather record.")
	logger.Println("Received PUT request to update weather record.")

	params := mux.Vars(r)
	id := params["id"]

	var weather models.Weather
	if err := db.DB.First(&weather, id).Error; err != nil {
		httpError(w, "Record not found", err, http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&weather)
	if err != nil {
		httpError(w, "Error decoding request body", err, http.StatusBadRequest)
		return
	}

	// Convert temperature from Celsius to Fahrenheit.
	weather.TemperatureF = utils.CelsiusToFahrenheit(weather.TemperatureF)

	if err := db.DB.Save(&weather).Error; err != nil {
		httpError(w, "Error updating record", err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, weather, http.StatusOK)
}

func DeleteWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received DELETE request for weather record.")
	logger.Println("Received DELETE request for weather record.")

	params := mux.Vars(r)
	id := params["id"]

	var weather models.Weather
	if err := db.DB.Delete(&weather, id).Error; err != nil {
		httpError(w, "Error deleting record", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SetupRoutes sets up the API routes using the gorilla/mux router.
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/weather", CreateWeather).Methods("POST")
	r.HandleFunc("/api/weather", GetWeather).Methods("GET")
	r.HandleFunc("/api/weather/{id}", GetWeatherByID).Methods("GET")
	r.HandleFunc("/api/weather/{id}", UpdateWeather).Methods("PUT")
	r.HandleFunc("/api/weather/{id}", DeleteWeather).Methods("DELETE")
	return r
}

// Utility function to send JSON response.
func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Utility function to handle HTTP errors.
func httpError(w http.ResponseWriter, message string, err error, statusCode int) {
	if err != nil {
		log.Println(message+":", err)
		logger.Println(message+":", err)
	} else {
		log.Println(message)
		logger.Println(message)
	}
	utils.HttpError(w, message, err, statusCode)
}
