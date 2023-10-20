package db

import (
	"database/sql"
	"log"
	"weather-history-api/configs"
	"weather-history-api/models"

	_ "github.com/lib/pq" // Importing the PostgreSQL driver.
)

var db *sql.DB

// InitDB initializes the database connection.
func InitDB(config configs.Config) {
	// Configure your database connection here.
	connStr := "user=" + config.DBUser + " dbname=" + config.DBName + " sslmode=disable password=" + config.DBPassword
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database successfully connected!")
}

// CloseDB closes the database connection.
func CloseDB() {
	if err := db.Close(); err != nil {
		log.Println("Error while closing database:", err)
	} else {
		log.Println("Database connection closed.")
	}
}

// InsertWeather inserts a weather record into the database and updates the ID field of the provided Weather object.
func InsertWeather(weather *models.Weather) error {
	// Prepare the SQL statement for inserting a new weather record.
	stmt, err := db.Prepare(`
        INSERT INTO weather (date, location, temperature)
        VALUES ($1, $2, $3)
        RETURNING id
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the weather data.
	err = stmt.QueryRow(weather.Date, weather.Location, weather.Temperature).Scan(&weather.ID)
	if err != nil {
		return err
	}

	return nil
}

// Additional database functions will be added here, such as FetchWeather, UpdateWeather, DeleteWeather, etc.
