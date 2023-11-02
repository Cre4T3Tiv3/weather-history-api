package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"weather-history-api/db"
	"weather-history-api/logger"
	"weather-history-api/models"
	"weather-history-api/utils"

	"github.com/gorilla/mux"
)

// CreateWeather handles the creation of weather records via HTTP POST request.
func CreateWeather(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Received POST request to /api/weather.")

	// Parse the request body.
	var weather models.Weather
	err := json.NewDecoder(r.Body).Decode(&weather)
	if err != nil {
		logger.Error.Println("Error decoding request body:", err)
		httpError(w, "Error decoding request body", err, http.StatusBadRequest)
		return
	}

	// Convert temperature from Celsius to Fahrenheit.
	weather.TemperatureF = utils.CelsiusToFahrenheit(weather.TemperatureF)

	logger.Info.Println("Parsed weather data:", weather)

	// Validate the weather data.
	if weather.Date.IsZero() || weather.Location == "" || weather.TemperatureF == 0 {
		logger.Error.Println("Invalid weather data.")
		httpError(w, "Invalid weather data", err, http.StatusBadRequest)
		return
	}

	// Insert the new weather record into the database.
	if err := db.DBConn.Create(&weather).Error; err != nil {
		if strings.Contains(err.Error(), "idx_weathers_location") {
			logger.Error.Println("Duplicate entry for location:", weather.Location)
			httpError(w, "Weather data already exists for this location.", err, http.StatusBadRequest)
		} else {
			logger.Error.Println("Error inserting into database:", err)
			httpError(w, "Error inserting into database:", err, http.StatusInternalServerError)
		}
		return
	}

	logger.Info.Println("Successfully inserted weather data into the database.")

	// Respond with the newly created weather record.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(weather)
	if err != nil {
		logger.Error.Println("Error encoding response:", err)
		httpError(w, "Error encoding response", err, http.StatusInternalServerError)
		return
	}
	jsonResponse(w, weather, http.StatusCreated)
	logger.Info.Println("Response sent successfully.")
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Received GET request for all weather records.")

	var weathers []models.Weather
	if err := db.DBConn.Find(&weathers).Error; err != nil {
		httpError(w, "Error fetching weather records", err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, weathers, http.StatusOK)
}

func GetWeatherByID(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Received GET request for a specific weather record.")

	params := mux.Vars(r)
	id := params["id"]

	var weather models.Weather
	if err := db.DBConn.First(&weather, id).Error; err != nil {
		httpError(w, "Record not found", err, http.StatusNotFound)
		return
	}

	jsonResponse(w, weather, http.StatusOK)
}

func UpdateWeather(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Received PUT request to update weather record.")

	params := mux.Vars(r)
	id := params["id"]

	var weather models.Weather
	if err := db.DBConn.First(&weather, id).Error; err != nil {
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

	if err := db.DBConn.Save(&weather).Error; err != nil {
		httpError(w, "Error updating record", err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, weather, http.StatusOK)
}

func DeleteWeather(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Received DELETE request for weather record.")

	params := mux.Vars(r)
	id := params["id"]

	var weather models.Weather
	if err := db.DBConn.Delete(&weather, id).Error; err != nil {
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
		logger.Error.Println(message+":", err)
	} else {
		logger.Info.Println(message)
	}
	utils.HttpError(w, message, err, statusCode)
}
