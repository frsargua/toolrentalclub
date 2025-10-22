package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/yourusername/toolrentalclub/domain/auth"
)

// AuthService implements the auth.Service interface using Firebase
type AuthService struct {
	app *firebase.App
}

// NewAuthService creates a new Firebase auth service
func NewAuthService(app *firebase.App) *AuthService {
	return &AuthService{
		app: app,
	}
}

// VerifyToken verifies a Firebase ID token and returns token information
func (s *AuthService) VerifyToken(ctx context.Context, tokenValue string) (*auth.Token, error) {
	if s.app == nil {
		return nil, fmt.Errorf("firebase app not initialized")
	}

	// Get Firebase Auth client
	client, err := s.app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth client: %w", err)
	}

	// Verify the token
	firebaseToken, err := client.VerifyIDToken(ctx, tokenValue)
	if err != nil {
		return nil, fmt.Errorf("invalid or expired token: %w", err)
	}

	// Extract email from claims
	email := ""
	if emailClaim, ok := firebaseToken.Claims["email"].(string); ok {
		email = emailClaim
	}

	// Create domain token
	token := auth.NewToken(tokenValue, firebaseToken.UID, email)
	return token, nil
}

