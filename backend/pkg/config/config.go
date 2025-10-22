package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Port                     string
	FirebaseCredentialsJSON  string
	FirebaseServiceAccount   string
}

// Load loads the configuration from environment variables
func Load() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:                    port,
		FirebaseCredentialsJSON: os.Getenv("FIREBASE_CREDENTIALS_JSON"),
		FirebaseServiceAccount:  os.Getenv("FIREBASE_SERVICE_ACCOUNT"),
	}
}

