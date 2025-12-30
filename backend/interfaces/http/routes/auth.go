package routes

import "github.com/gorilla/mux"

// registerAuthRoutes sets up all authentication-related endpoints
// These routes handle token verification and other auth operations
func (rt *Router) registerAuthRoutes(r *mux.Router) {
	authRouter := r.PathPrefix("/api/auth").Subrouter()

	// POST /api/auth/verify - Verify Firebase token
	authRouter.HandleFunc("/verify", rt.authHandler.VerifyToken).Methods("POST")
}
