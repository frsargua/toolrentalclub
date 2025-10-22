package auth

import "context"

// Service defines the interface for authentication operations
type Service interface {
	// VerifyToken verifies an authentication token and returns token information
	VerifyToken(ctx context.Context, token string) (*Token, error)
}

