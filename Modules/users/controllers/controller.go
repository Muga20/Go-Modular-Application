package controllers

import (
	"net/http"
)

// UsersController handles users-related requests
type UsersController struct {}

// NewUsersController creates a new UsersController
func NewUsersController() *UsersController {
	return &UsersController{}
}

// GetHandler handles GET requests
func (c *UsersController) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users Get Handler"))
}
