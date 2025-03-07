package controllers

import (
	"net/http"
)

// AuthController handles auth-related requests
type AuthController struct {}

// NewAuthController creates a new AuthController
func NewAuthController() *AuthController {
	return &AuthController{}
}

// GetHandler handles GET requests
func (c *AuthController) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Auth Get Handler"))
}
