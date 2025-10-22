package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/toolrentalclub/interfaces/http/dto"
)

// respondWithJSON writes a JSON response with the given status code
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError writes a JSON error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, dto.ErrorResponse{Error: message})
}

