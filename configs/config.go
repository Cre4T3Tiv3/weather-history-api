// configs/config.go
package configs

// Config represents the configuration structure for the application.
type Config struct {
	DBUser     string // DBUser is the username for the database connection.
	DBName     string // DBName is the name of the database.
	DBPassword string // DBPassword is the password for the database connection.
	// ... any other fields that may be added in the future
}
