package logger

import (
	"io"
	"log"
	"os"
)

var (
	// Info handles informational messages.
	Info *log.Logger
	// Error handles error-level messages.
	Error *log.Logger
)

// InitLog initializes the loggers.
func InitLog() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	// Using a multi-writer to log to both file and console.
	multiWriterInfo := io.MultiWriter(file, os.Stdout)
	multiWriterError := io.MultiWriter(file, os.Stderr)

	Info = log.New(multiWriterInfo, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(multiWriterError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
