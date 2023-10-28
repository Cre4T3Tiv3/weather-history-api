package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"weather-history-api/db"
	"weather-history-api/logger"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ensureCorrectDirectory(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)             // Get current file's path.
	dir := filepath.Join(filepath.Dir(filename), "..") // Move up one directory to `weather-history-api`.
	err := os.Chdir(dir)
	if err != nil && t != nil {
		t.Fatalf("Failed to change to correct working directory: %s", err)
	}
}

func TestWorkingDirectory(t *testing.T) {
	ensureCorrectDirectory(t)

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %s", err)
	}
	t.Logf("Current working directory: %s", dir)
}

func TestMain(m *testing.M) {
	ensureCorrectDirectory(nil) // Use nil since it's not a standard test.

	absPath, _ := filepath.Abs("logs/app.log")
	fmt.Println("Looking for log file at:", absPath)

	if _, err := os.Stat("logs/app.log"); os.IsNotExist(err) {
		fmt.Println("Log file does not exist!")
	} else if err != nil {
		fmt.Println("Error checking log file:", err)
	} else {
		fmt.Println("Log file exists!")
	}

	// Initialize the logger.
	logger.InitLogger()

	// Run all the tests.
	code := m.Run()

	// Exit with the code returned from the tests.
	os.Exit(code)
}

func TestGetWeatherHandler(t *testing.T) {
	ensureCorrectDirectory(t)

	// Backup the original DB connection for later restoration.
	originalDB := db.GetDBConnection()

	// Set up mock database.
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %s", err)
	}
	defer mockDB.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open GORM DB: %s", err)
	}

	// Use the mock DB for the test.
	db.SetDBConnection(gormDB)

	rows := sqlmock.NewRows([]string{"ID", "City", "Temperature", "Condition"}).
		AddRow(1, "London", 20, "Sunny").
		AddRow(2, "Paris", 25, "Cloudy")
	mock.ExpectQuery("^SELECT (.+) FROM weather$").WillReturnRows(rows)

	// Check if there were any issues with the expectations.
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Restore the original DB connection after the test.
	db.SetDBConnection(originalDB)
}
