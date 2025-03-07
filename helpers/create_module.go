package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ModuleTemplate represents the data needed to generate a module
type ModuleTemplate struct {
	ModuleName string
}

func main() {
	// Get the module name from the command line
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run helpers/create_module.go <module_name>")
		return
	}
	moduleName := os.Args[1]

	// Define the module structure
	moduleDir := filepath.Join("Modules", moduleName)
	dirs := []string{
		filepath.Join(moduleDir, "controllers"),
		filepath.Join(moduleDir, "models"),
		filepath.Join(moduleDir, "routes"),
		filepath.Join(moduleDir, "middleware"),
		filepath.Join(moduleDir, "utils"),
		filepath.Join(moduleDir, "config"),
		filepath.Join(moduleDir, "tests"),
		filepath.Join(moduleDir, "migrations"), // Add migrations directory
	}

	// Create directories
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", dir, err)
			return
		}
	}

	// Define template data
	data := ModuleTemplate{
		ModuleName: moduleName,
	}

	// Create files from templates
	files := []struct {
		path     string
		template string
	}{
		{filepath.Join(moduleDir, "controllers", "controller.go"), controllerTemplate},
		{filepath.Join(moduleDir, "models", "model.go"), modelTemplate},
		{filepath.Join(moduleDir, "routes", "routes.go"), routesTemplate},
		{filepath.Join(moduleDir, "middleware", "middleware.go"), middlewareTemplate},
		{filepath.Join(moduleDir, "utils", "utils.go"), utilsTemplate},
		{filepath.Join(moduleDir, "config", "config.go"), configTemplate},
		{filepath.Join(moduleDir, "tests", "controller_test.go"), testTemplate},
		{filepath.Join(moduleDir, "migrations", "0001_initial_migration.sql"), migrationTemplate}, // Add migration file
		{filepath.Join(moduleDir, "migrate.go"), migrateTemplate},                                 // Add migrate.go file
		{filepath.Join(moduleDir, "README.md"), readmeTemplate},
	}

	for _, file := range files {
		if err := createFileFromTemplate(file.path, file.template, data); err != nil {
			fmt.Printf("Failed to create file %s: %v\n", file.path, err)
			return
		}
	}

	fmt.Printf("Module '%s' created successfully!\n", moduleName)
}

// createFileFromTemplate creates a file from a template
func createFileFromTemplate(path, tmpl string, data ModuleTemplate) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Add a custom function to capitalize the first letter
	funcMap := template.FuncMap{
		"Title": strings.Title,
	}

	t := template.Must(template.New("").Funcs(funcMap).Parse(tmpl))
	return t.Execute(file, data)
}

// Templates for module files
var (
	controllerTemplate = `package controllers

import (
	"net/http"
)

// {{.ModuleName | Title}}Controller handles {{.ModuleName}}-related requests
type {{.ModuleName | Title}}Controller struct {}

// New{{.ModuleName | Title}}Controller creates a new {{.ModuleName | Title}}Controller
func New{{.ModuleName | Title}}Controller() *{{.ModuleName | Title}}Controller {
	return &{{.ModuleName | Title}}Controller{}
}

// GetHandler handles GET requests
func (c *{{.ModuleName | Title}}Controller) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{{.ModuleName | Title}} Get Handler"))
}
`

	modelTemplate = `package models

// {{.ModuleName | Title}} represents a {{.ModuleName}} entity
type {{.ModuleName | Title}} struct {
	ID   int
	Name string
}
`

	routesTemplate = `package routes

import (
	"net/http"
	"auto_verse/Modules/{{.ModuleName}}/controllers"
	"auto_verse/Modules/{{.ModuleName}}/middleware"
)

// Setup{{.ModuleName | Title}}Routes configures routes for the {{.ModuleName}} module
func Setup{{.ModuleName | Title}}Routes() {
	controller := controllers.New{{.ModuleName | Title}}Controller()
	http.HandleFunc("/{{.ModuleName}}", middleware.LogRequest(controller.GetHandler))
}
`

	middlewareTemplate = `package middleware

import (
	"net/http"
	"log"
)

// LogRequest logs incoming HTTP requests
func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}
`

	utilsTemplate = `package utils

// Utility functions for the {{.ModuleName}} module
func {{.ModuleName | Title}}Utility() string {
	return "{{.ModuleName | Title}} Utility"
}
`

	configTemplate = `package config

// {{.ModuleName | Title}}Config holds configuration for the {{.ModuleName}} module
type {{.ModuleName | Title}}Config struct {
	Enabled bool
}
`

	testTemplate = `package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"auto_verse/Modules/{{.ModuleName}}/controllers"
)

func Test{{.ModuleName | Title}}Controller_GetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/{{.ModuleName}}", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	controller := controllers.New{{.ModuleName | Title}}Controller()
	controller.GetHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := "{{.ModuleName | Title}} Get Handler"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}
`

	migrationTemplate = `-- 0001_initial_migration.sql
-- Add your SQL statements here
CREATE TABLE IF NOT EXISTS {{.ModuleName}} (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

	migrateTemplate = `package {{.ModuleName}}

import (
	"database/sql"
	"auto_verse/migrations"
)

func init() {
	// Register the {{.ModuleName}} migration function
	migrations.Register(Migrate)
}

// Migrate runs all migrations for the {{.ModuleName}} module
func Migrate(db *sql.DB) error {
	// Migration logic for the {{.ModuleName}} module
	return nil
}
`

	readmeTemplate = `# {{.ModuleName | Title}} Module

This module handles {{.ModuleName}}-related functionality.

## Structure
- **controllers/**: Handles HTTP requests.
- **models/**: Defines database models.
- **routes/**: Defines API routes.
- **middleware/**: Middleware for the module.
- **utils/**: Utility functions.
- **config/**: Module-specific configuration.
- **tests/**: Unit and integration tests.
- **migrations/**: Database migration files.

## Usage
- To use the module, import it in your application and call the setup function in your main file.

## Migrations
- Run migrations using the ` + "`Migrate`" + ` function in ` + "`migrate.go`" + `.

## Testing
- Run ` + "`go test ./Modules/{{.ModuleName}}/tests`" + ` to run tests for this module.
`
)
