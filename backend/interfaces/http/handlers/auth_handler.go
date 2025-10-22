package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/toolrentalclub/application/auth"
	"github.com/yourusername/toolrentalclub/interfaces/http/dto"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	authUseCase *auth.UseCase
}

// NewAuthHandler creates a new authentication handler
func NewAuthHandler(authUseCase *auth.UseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// VerifyToken handles token verification requests
func (h *AuthHandler) VerifyToken(w http.ResponseWriter, r *http.Request) {
	var req dto.VerifyTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Token == "" {
		respondWithError(w, http.StatusBadRequest, "Token is required")
		return
	}

	// Verify token and get or create user
	token, user, err := h.authUseCase.VerifyTokenAndGetUser(r.Context(), req.Token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	response := dto.VerifyTokenResponse{
		Success: true,
		Message: "Token verified successfully",
		UserID:  user.ID,
		Email:   token.Email,
	}

	respondWithJSON(w, http.StatusOK, response)
}

