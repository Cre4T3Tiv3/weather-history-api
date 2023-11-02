// Utility function to handle HTTP errors.

package utils

import (
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

// Initialize the logger. This should be called during application startup.
func InitLogger() {
	logger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func HttpError(w http.ResponseWriter, message string, err error, statusCode int) {
	errorMsg := message
	if err != nil {
		errorMsg += ": " + err.Error()
	}
	// Use 'errorMsg' as the message to log and send as the HTTP response.
	if logger != nil {
		logger.Println(errorMsg)
	} else {
		log.Println("Logger not initialized. Error:", errorMsg)
	}
	http.Error(w, errorMsg, statusCode)
}
