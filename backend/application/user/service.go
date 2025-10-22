package user

import (
	"context"

	"github.com/yourusername/toolrentalclub/domain/user"
)

// UseCase represents the user use cases
type UseCase struct {
	userRepo user.Repository
}

// NewUseCase creates a new user use case
func NewUseCase(userRepo user.Repository) *UseCase {
	return &UseCase{
		userRepo: userRepo,
	}
}

// GetUserByID retrieves a user by their ID
func (uc *UseCase) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	return uc.userRepo.FindByID(ctx, id)
}

// GetUserByEmail retrieves a user by their email
func (uc *UseCase) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	return uc.userRepo.FindByEmail(ctx, email)
}

