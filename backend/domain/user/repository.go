package user

import "context"

// Repository defines the interface for user data operations
type Repository interface {
	// FindByID retrieves a user by their ID
	FindByID(ctx context.Context, id string) (*User, error)
	
	// FindByEmail retrieves a user by their email
	FindByEmail(ctx context.Context, email string) (*User, error)
	
	// Create creates a new user
	Create(ctx context.Context, user *User) error
	
	// Update updates an existing user
	Update(ctx context.Context, user *User) error
}

