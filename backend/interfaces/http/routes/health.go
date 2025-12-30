package routes

import "github.com/gorilla/mux"

// registerHealthRoutes sets up all health-related endpoints
// These routes are public and do not require authentication
func (rt *Router) registerHealthRoutes(r *mux.Router) {
	// GET /api/health - Health check endpoint
	r.HandleFunc("/api/health", rt.healthHandler.HealthCheck).Methods("GET")
}
