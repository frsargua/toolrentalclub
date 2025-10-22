package auth

import (
	"context"

	"github.com/yourusername/toolrentalclub/domain/auth"
	"github.com/yourusername/toolrentalclub/domain/user"
)

// UseCase represents the authentication use cases
type UseCase struct {
	authService auth.Service
	userRepo    user.Repository
}

// NewUseCase creates a new authentication use case
func NewUseCase(authService auth.Service, userRepo user.Repository) *UseCase {
	return &UseCase{
		authService: authService,
		userRepo:    userRepo,
	}
}

// VerifyTokenAndGetUser verifies a token and returns or creates the user
func (uc *UseCase) VerifyTokenAndGetUser(ctx context.Context, tokenValue string) (*auth.Token, *user.User, error) {
	// Verify the token
	token, err := uc.authService.VerifyToken(ctx, tokenValue)
	if err != nil {
		return nil, nil, err
	}

	// Try to find the user
	existingUser, err := uc.userRepo.FindByID(ctx, token.UserID)
	if err == nil && existingUser != nil {
		return token, existingUser, nil
	}

	// If user doesn't exist, create a new one
	newUser := user.NewUser(token.UserID, token.Email)
	if err := uc.userRepo.Create(ctx, newUser); err != nil {
		// If creation fails, it might be a race condition, try to find again
		existingUser, findErr := uc.userRepo.FindByID(ctx, token.UserID)
		if findErr != nil {
			return nil, nil, err
		}
		return token, existingUser, nil
	}

	return token, newUser, nil
}

// VerifyToken verifies a token without user operations
func (uc *UseCase) VerifyToken(ctx context.Context, tokenValue string) (*auth.Token, error) {
	return uc.authService.VerifyToken(ctx, tokenValue)
}

