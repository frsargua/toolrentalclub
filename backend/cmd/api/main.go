package main

import (
	"context"
	"log"
	"time"

	authApp "github.com/yourusername/toolrentalclub/application/auth"
	userApp "github.com/yourusername/toolrentalclub/application/user"
	"github.com/yourusername/toolrentalclub/infrastructure/firebase"
	"github.com/yourusername/toolrentalclub/infrastructure/repository/memory"
	"github.com/yourusername/toolrentalclub/interfaces/http/handlers"
	"github.com/yourusername/toolrentalclub/interfaces/http/routes"
	"github.com/yourusername/toolrentalclub/pkg/config"
	"github.com/yourusername/toolrentalclub/pkg/server"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize Firebase
	ctx := context.Background()
	firebaseApp, err := firebase.InitializeApp(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	// Initialize repositories
	userRepo := memory.NewUserRepository()

	// Initialize domain services
	var authService *firebase.AuthService
	if firebaseApp != nil {
		authService = firebase.NewAuthService(firebaseApp)
	} else {
		log.Println("WARNING: Firebase not initialized. Authentication will not work.")
	}

	// Initialize application use cases
	authUseCase := authApp.NewUseCase(authService, userRepo)
	userUseCase := userApp.NewUseCase(userRepo)

	// Initialize HTTP handlers
	healthHandler := handlers.NewHealthHandler()
	authHandler := handlers.NewAuthHandler(authUseCase)
	userHandler := handlers.NewUserHandler(userUseCase)

	// Setup router with all routes
	router := routes.NewRouter(
		healthHandler,
		authHandler,
		userHandler,
		authUseCase,
		authService != nil,
	)
	r := router.Setup()

	// Start server
	serverCfg := server.Config{
		Port:         cfg.Port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(server.Start(r, serverCfg))
}
