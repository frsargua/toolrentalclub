package dto

// VerifyTokenRequest represents the request to verify a token
type VerifyTokenRequest struct {
	Token string `json:"token"`
}

// VerifyTokenResponse represents the response from token verification
type VerifyTokenResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UserID  string `json:"userId,omitempty"`
	Email   string `json:"email,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

