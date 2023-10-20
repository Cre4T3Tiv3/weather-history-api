package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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
	if err := db.InsertWeather(&weather); err != nil {
		log.Println("Error inserting into database:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	log.Println("Response sent successfully.")
}

// GetWeather handles HTTP GET requests for weather records.
func GetWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request for weather records.")
	// Temporary response for stubbed-out handler.
	w.Write([]byte("GET handler not implemented yet."))
}

// UpdateWeather handles HTTP PUT requests to update weather records.
func UpdateWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received PUT request to update weather record.")
	// Temporary response for stubbed-out handler.
	w.Write([]byte("PUT handler not implemented yet."))
}

// DeleteWeather handles HTTP DELETE requests to delete weather records.
func DeleteWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Received DELETE request for weather record.")
	// Temporary response for stubbed-out handler.
	w.Write([]byte("DELETE handler not implemented yet."))
}

// SetupRoutes configures and returns the router with API routes.
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/weather", CreateWeather).Methods("POST")
	r.HandleFunc("/api/weather", GetWeather).Methods("GET")
	r.HandleFunc("/api/weather/{id}", UpdateWeather).Methods("PUT")
	r.HandleFunc("/api/weather/{id}", DeleteWeather).Methods("DELETE")
	return r
}
