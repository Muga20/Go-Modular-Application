package routes

import (
	"auto_verse/Modules/users/controllers"
	"auto_verse/Modules/users/middleware"
	"net/http"
)

// SetupUsersRoutes configures routes for the users module and wraps them under /api/v1
func SetupUsersRoutes(router *http.ServeMux) {
	controller := controllers.NewUsersController()

	// Create a sub-router for /api/v1
	apiRouter := http.NewServeMux()

	// Register routes under /api/v1/users
	apiRouter.HandleFunc("/users", middleware.LogRequest(controller.GetHandler))

	// Wrap the sub-router under /api/v1
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))
}
