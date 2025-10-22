package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/yourusername/toolrentalclub/application/auth"
)

// AuthMiddleware creates middleware that validates authentication tokens
func AuthMiddleware(authUseCase *auth.UseCase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the token from the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				respondWithError(w, http.StatusUnauthorized, "Authorization header required")
				return
			}

			// Extract the token (format: "Bearer <token>")
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				respondWithError(w, http.StatusUnauthorized, "Invalid authorization header format")
				return
			}

			idToken := parts[1]

			// Verify the token
			token, err := authUseCase.VerifyToken(r.Context(), idToken)
			if err != nil {
				respondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
				return
			}

			// Add user info to context
			ctx := context.WithValue(r.Context(), "userID", token.UserID)
			ctx = context.WithValue(ctx, "email", token.Email)

			// Call the next handler with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// respondWithError is a helper function to send error responses
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// Simple error response without importing the handlers package
	w.Write([]byte(`{"error":"` + message + `"}`))
}

