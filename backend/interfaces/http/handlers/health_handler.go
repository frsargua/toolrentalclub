package handlers

import (
	"net/http"

	"github.com/yourusername/toolrentalclub/interfaces/http/dto"
)

// HealthHandler handles health check requests
type HealthHandler struct{}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck handles health check requests
func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := dto.HealthCheckResponse{
		Status:  "ok",
		Message: "Tool Rental Club API is running",
	}
	respondWithJSON(w, http.StatusOK, response)
}

