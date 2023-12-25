package utils

import (
	"encoding/json"
	"net/http"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func LogErrorf(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Fatal(errMsg)
}

func HandleError(w http.ResponseWriter, message string, statusCode int) {
	errorResponse := ErrorResponse{
		Tag:   http.StatusText(statusCode),
		Value: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse)
}

// Reponse format error message
