package models

import (
	"time"
)

type Weather struct {
	ID          uint      `gorm:"primaryKey"` // Define ID as the primary key
	Date        time.Time `json:"date"`
	Location    string    `gorm:"uniqueIndex;not null"` // Define a unique index on Location
	Temperature float64   `json:"temperature"`
	CreatedAt   time.Time // These fields are automatically populated by GORM
	UpdatedAt   time.Time
}
