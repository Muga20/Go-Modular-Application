package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// MigrationFunc is a function that runs migrations for a module
type MigrationFunc func(db *sql.DB) error

var (
	registry []MigrationFunc
	mu       sync.Mutex
)

// Register adds a migration function to the registry
func Register(migrationFunc MigrationFunc) {
	mu.Lock()
	defer mu.Unlock()
	registry = append(registry, migrationFunc)
}

// RunAll runs all registered migrations
func RunAll(db *sql.DB) error {
	// Get a list of all modules
	modulesDir := "Modules"
	modules, err := os.ReadDir(modulesDir)
	if err != nil {
		return fmt.Errorf("failed to read modules directory: %v", err)
	}

	// Iterate through each module and apply its migrations
	for _, module := range modules {
		if module.IsDir() {
			moduleName := module.Name()
			migrationsDir := filepath.Join(modulesDir, moduleName, "migrations")

			// Check if the migrations directory exists
			if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
				log.Printf("No migrations directory found for module: %s. Skipping...", moduleName)
				continue
			}

			// Check if there are any migration files
			migrationFiles, err := filepath.Glob(filepath.Join(migrationsDir, "*.up.sql"))
			if err != nil {
				return fmt.Errorf("failed to list migration files for module %s: %v", moduleName, err)
			}
			if len(migrationFiles) == 0 {
				log.Printf("No migration files found for module: %s. Skipping...", moduleName)
				continue
			}

			// Apply migrations for this module
			if err := applyMigrations(db, migrationsDir); err != nil {
				return fmt.Errorf("failed to apply migrations for module %s: %v", moduleName, err)
			}

			log.Printf("Migrations applied successfully for module: %s", moduleName)
		}
	}

	log.Println("All migrations applied successfully!")
	return nil
}

// RunForModule runs migrations for a specific module
func RunForModule(db *sql.DB, moduleName, direction string) error {
	migrationsDir := filepath.Join("Modules", moduleName, "migrations")

	// Check if the migrations directory exists
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return fmt.Errorf("no migrations directory found for module: %s", moduleName)
	}

	// Check if there are any migration files
	migrationFiles, err := filepath.Glob(filepath.Join(migrationsDir, "*."+direction+".sql"))
	if err != nil {
		return fmt.Errorf("failed to list migration files for module %s: %v", moduleName, err)
	}
	if len(migrationFiles) == 0 {
		return fmt.Errorf("no migration files found for module: %s", moduleName)
	}

	// Apply or rollback migrations for this module
	switch direction {
	case "up":
		if err := applyMigrations(db, migrationsDir); err != nil {
			return fmt.Errorf("failed to apply migrations for module %s: %v", moduleName, err)
		}
	case "down":
		if err := rollbackMigrations(db, migrationsDir); err != nil {
			return fmt.Errorf("failed to rollback migrations for module %s: %v", moduleName, err)
		}
	default:
		return fmt.Errorf("invalid migration direction: %s", direction)
	}

	log.Printf("Migrations %s applied successfully for module: %s", direction, moduleName)
	return nil
}

// RollbackAll rolls back all migrations (down)
func RollbackAll(db *sql.DB) error {
	// Get a list of all modules
	modulesDir := "Modules"
	modules, err := os.ReadDir(modulesDir)
	if err != nil {
		return fmt.Errorf("failed to read modules directory: %v", err)
	}

	// Iterate through each module and rollback its migrations
	for _, module := range modules {
		if module.IsDir() {
			moduleName := module.Name()
			migrationsDir := filepath.Join(modulesDir, moduleName, "migrations")

			// Check if the migrations directory exists
			if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
				log.Printf("No migrations directory found for module: %s. Skipping...", moduleName)
				continue
			}

			// Check if there are any migration files
			migrationFiles, err := filepath.Glob(filepath.Join(migrationsDir, "*.down.sql"))
			if err != nil {
				return fmt.Errorf("failed to list migration files for module %s: %v", moduleName, err)
			}
			if len(migrationFiles) == 0 {
				log.Printf("No migration files found for module: %s. Skipping...", moduleName)
				continue
			}

			// Rollback migrations for this module
			if err := rollbackMigrations(db, migrationsDir); err != nil {
				return fmt.Errorf("failed to rollback migrations for module %s: %v", moduleName, err)
			}

			log.Printf("Migrations rolled back successfully for module: %s", moduleName)
		}
	}

	log.Println("All migrations rolled back successfully!")
	return nil
}

// applyMigrations applies migrations for a specific module
func applyMigrations(db *sql.DB, migrationsDir string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsDir, // Path to migration files
		"mysql",                 // Database driver
		driver,                  // Database instance
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations (up): %v", err)
	}

	return nil
}

// rollbackMigrations rolls back migrations for a specific module
func rollbackMigrations(db *sql.DB, migrationsDir string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsDir, // Path to migration files
		"mysql",                 // Database driver
		driver,                  // Database instance
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate instance: %v", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migrations (down): %v", err)
	}

	return nil
}
