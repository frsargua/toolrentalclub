package firebase

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// InitializeApp initializes and returns a Firebase app
func InitializeApp(ctx context.Context) (*firebase.App, error) {
	// Try to get Firebase credentials from environment
	credentialsB64 := os.Getenv("FIREBASE_CREDENTIALS_JSON")
	credentialsPath := os.Getenv("FIREBASE_SERVICE_ACCOUNT")

	if credentialsB64 != "" {
		// Decode base64 credentials
		credentialsJSON, err := base64.StdEncoding.DecodeString(credentialsB64)
		if err != nil {
			return nil, fmt.Errorf("error decoding Firebase credentials: %w", err)
		}

		// Use decoded credentials
		opt := option.WithCredentialsJSON(credentialsJSON)
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			return nil, fmt.Errorf("error initializing firebase app from JSON: %w", err)
		}
		log.Println("Firebase initialized successfully from environment variable (base64 decoded)")
		return app, nil
	} else if credentialsPath != "" {
		// Use credentials from file path
		opt := option.WithCredentialsFile(credentialsPath)
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			return nil, fmt.Errorf("error initializing firebase app from file: %w", err)
		}
		log.Println("Firebase initialized successfully from file")
		return app, nil
	}

	log.Println("WARNING: Neither FIREBASE_CREDENTIALS_JSON nor FIREBASE_SERVICE_ACCOUNT set. Firebase features will not work.")
	return nil, nil
}

