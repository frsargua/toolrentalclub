package dto

// UserProfileResponse represents a user profile response
type UserProfileResponse struct {
	UserID  string `json:"userId"`
	Email   string `json:"email"`
	Message string `json:"message,omitempty"`
}

// HealthCheckResponse represents a health check response
type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

