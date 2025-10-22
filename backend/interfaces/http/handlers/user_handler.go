package handlers

import (
	"net/http"

	"github.com/yourusername/toolrentalclub/application/user"
	"github.com/yourusername/toolrentalclub/interfaces/http/dto"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userUseCase *user.UseCase
}

// NewUserHandler creates a new user handler
func NewUserHandler(userUseCase *user.UseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// GetProfile handles requests to get the authenticated user's profile
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok || userID == "" {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized - authentication required")
		return
	}

	// Get user from repository
	user, err := h.userUseCase.GetUserByID(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	response := dto.UserProfileResponse{
		UserID:  user.ID,
		Email:   user.Email,
		Message: "This is a protected route",
	}

	respondWithJSON(w, http.StatusOK, response)
}

