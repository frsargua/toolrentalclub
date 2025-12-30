package routes

import (
	"github.com/gorilla/mux"

	authApp "github.com/yourusername/toolrentalclub/application/auth"
	"github.com/yourusername/toolrentalclub/interfaces/http/handlers"
	"github.com/yourusername/toolrentalclub/interfaces/http/middleware"
)

// Router holds all the dependencies needed for route registration
type Router struct {
	healthHandler *handlers.HealthHandler
	authHandler   *handlers.AuthHandler
	userHandler   *handlers.UserHandler
	authUseCase   *authApp.UseCase
	authEnabled   bool
}

// NewRouter creates a new Router with all required dependencies
func NewRouter(
	healthHandler *handlers.HealthHandler,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	authUseCase *authApp.UseCase,
	authEnabled bool,
) *Router {
	return &Router{
		healthHandler: healthHandler,
		authHandler:   authHandler,
		userHandler:   userHandler,
		authUseCase:   authUseCase,
		authEnabled:   authEnabled,
	}
}

// Setup creates and configures the main router with all routes and middleware
func (rt *Router) Setup() *mux.Router {
	r := mux.NewRouter()

	// Apply global middleware
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.LoggingMiddleware)

	// Register all route groups
	rt.registerHealthRoutes(r)
	rt.registerAuthRoutes(r)
	rt.registerProtectedRoutes(r)

	return r
}
