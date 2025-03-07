package routes

import (
	"net/http"
	"auto_verse/Modules/auth/controllers"
	"auto_verse/Modules/auth/middleware"
)

// SetupAuthRoutes configures routes for the auth module
func SetupAuthRoutes() {
	controller := controllers.NewAuthController()
	http.HandleFunc("/auth", middleware.LogRequest(controller.GetHandler))
}
