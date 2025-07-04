package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chwnsng/Guessing-Game/backend/models"
)

// sending responses with status code and JSON payload
// accepts any type to alow for different response bodies
func RespondJSON(w http.ResponseWriter, httpCode int, payload interface{}) {
	response, err := json.Marshal(payload)

	// send status 500 if failed to marshal payload
	if err != nil {
		log.Printf("Error marshalling payload to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}

// sending errors
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, models.ErrorResponse{Message: message})
}
