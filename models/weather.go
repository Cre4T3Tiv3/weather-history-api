package models

import (
	"time"
)

// Weather represents a weather data record.
type Weather struct {
	ID          uint      `gorm:"primaryKey"`           // Define ID as the primary key.
	Date        time.Time `json:"date"`                 // Date of the weather record.
	Location    string    `gorm:"uniqueIndex;not null"` // Location where the weather data was recorded (with a unique index).
	Temperature float64   `json:"temperature"`          // Temperature in Fahrenheit.
	CreatedAt   time.Time // Automatically populated timestamp for record creation.
	UpdatedAt   time.Time // Automatically populated timestamp for record update.
}
