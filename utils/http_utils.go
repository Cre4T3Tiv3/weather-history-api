// Utility function to handle HTTP errors.

package utils

import (
	"log"
	"net/http"
)

var logger *log.Logger

func HttpError(w http.ResponseWriter, message string, err error, statusCode int) {
	errorMsg := message
	if err != nil {
		errorMsg += ": " + err.Error()
	}
	// Use 'errorMsg' as the message to log and send as the HTTP response.
	log.Println(errorMsg)
	logger.Println(errorMsg)
	http.Error(w, errorMsg, statusCode)
}
