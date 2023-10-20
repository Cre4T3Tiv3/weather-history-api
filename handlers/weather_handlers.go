package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"weather-history-api/db"
	"weather-history-api/models"

	"github.com/gorilla/mux"
)

// CreateWeather handles the creation of weather records via HTTP POST request.
func CreateWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST request to /api/weather.")

	// Parse the request body.
	var weather models.Weather
	err := json.NewDecoder(r.Body).Decode(&weather)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Parsed weather data:", weather)

	// Validate the weather data (you can add more validation as needed).
	if weather.Date.IsZero() || weather.Location == "" || weather.Temperature == 0 {
		log.Println("Invalid weather data.")
		http.Error(w, "Invalid weather data", http.StatusBadRequest)
		return
	}

	// Insert the new weather record into the database.
	if err := db.DB.Create(&weather).Error; err != nil {
		if strings.Contains(err.Error(), "idx_weathers_location") {
			log.Println("Duplicate entry for location:", weather.Location)
			http.Error(w, "Weather data already exists for this location.", http.StatusBadRequest)
		} else {
			log.Println("Error inserting into database:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	log.Println("Successfully inserted weather data into the database.")

	// Respond with the newly created weather record.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(weather) // Note the = instead of :=
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	jsonResponse(w, weather, http.StatusCreated)
	log.Println("Response sent successfully.")
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request for weather records.")

	var weathers []models.Weather
	if err := db.DB.Find(&weathers).Error; err != nil {
		httpError(w, "Error fetching weather records", err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, weathers, http.StatusOK)
}

func UpdateWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received PUT request to update weather record.")

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

	if err := db.DB.Save(&weather).Error; err != nil {
		httpError(w, "Error updating record", err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, weather, http.StatusOK)
}

func DeleteWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received DELETE request for weather record.")

	params := mux.Vars(r)
	id := params["id"]

	var weather models.Weather
	if err := db.DB.Delete(&weather, id).Error; err != nil {
		httpError(w, "Error deleting record", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/weather", CreateWeather).Methods("POST")
	r.HandleFunc("/api/weather", GetWeather).Methods("GET")
	r.HandleFunc("/api/weather/{id}", UpdateWeather).Methods("PUT")
	r.HandleFunc("/api/weather/{id}", DeleteWeather).Methods("DELETE")
	return r
}

// Utility function to send JSON response
func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Utility function to handle HTTP errors
func httpError(w http.ResponseWriter, message string, err error, statusCode int) {
	if err != nil {
		log.Println(message+":", err)
	} else {
		log.Println(message)
	}
	http.Error(w, message, statusCode)
}
