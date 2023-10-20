// Utility function to handle HTTP errors.

package utils

import (
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	// Initialize the logger.
	logFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening or creating log file: %v", err)
	}
	// Create a logger instance to write logs to the file.
	logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
}

func HttpError(w http.ResponseWriter, message string, err error, statusCode int) {
	errorMsg := message
	if err != nil {
		errorMsg += ": " + err.Error()
	}
	// Now, use 'errorMsg' as the message to log and send as the HTTP response.
	log.Println(errorMsg)
	logger.Println(errorMsg)
	http.Error(w, errorMsg, statusCode)
}
