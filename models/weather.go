package models

import "time"

// Weather represents a weather data model.
type Weather struct {
	ID          int       `json:"id"`          // Unique identifier for the weather record.
	Date        time.Time `json:"date"`        // Date and time of the weather data.
	Location    string    `json:"location"`    // Location for which the weather data is recorded.
	Temperature float64   `json:"temperature"` // Temperature recorded at the location.
}
