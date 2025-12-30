package routes

import (
	"github.com/gorilla/mux"

	"github.com/yourusername/toolrentalclub/interfaces/http/middleware"
)

// registerProtectedRoutes sets up all routes that require authentication
// These routes are protected by the auth middleware
func (rt *Router) registerProtectedRoutes(r *mux.Router) {
	protectedRouter := r.PathPrefix("/api").Subrouter()

	// Apply authentication middleware if auth is enabled
	if rt.authEnabled {
		protectedRouter.Use(middleware.AuthMiddleware(rt.authUseCase))
	}

	// GET /api/profile - Get current user's profile
	protectedRouter.HandleFunc("/profile", rt.userHandler.GetProfile).Methods("GET")
}
