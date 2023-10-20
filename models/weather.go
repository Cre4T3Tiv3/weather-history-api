package models

import (
	"time"
)

// Weather represents a weather data record.
type Weather struct {
	ID           uint      `gorm:"primaryKey"`           // Define ID as the primary key.
	Date         time.Time `json:"date"`                 // Date of the weather record.
	Location     string    `gorm:"uniqueIndex;not null"` // Location where the weather data was recorded (with a unique index).
	TemperatureF float64   `json:"temperatureF"`         // Temperature in Fahrenheit.
	CreatedAt    time.Time // Automatically populated timestamp for record creation.
	UpdatedAt    time.Time // Automatically populated timestamp for record update.
}

func (w *Weather) ValidateTemperature() bool {
	// Assuming that it's unlikely to have temperatures below -100°F.
	// or above 150°F. Adjust this range as needed.
	if w.TemperatureF < -100 || w.TemperatureF > 150 {
		return false
	}
	return true
}
