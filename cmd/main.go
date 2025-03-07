package main

import (
	"auto_verse/Modules/users/routes" // Import the users routes package
	"auto_verse/config"
	"auto_verse/migrations"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var cfg = mysql.Config{
	User:                 config.Envs.DBUser,
	Passwd:               config.Envs.DBPassword,
	Addr:                 config.Envs.DBAddress,
	DBName:               config.Envs.DBName,
	Net:                  "tcp",
	AllowNativePasswords: true,
	ParseTime:            true,
}

func main() {
	// Parse command-line flags
	migrateCmd := flag.String("migrate", "", "Run migrations (up, down, or force <version>)")
	moduleName := flag.String("module", "", "Specify the module to run migrations for (e.g., users, auth)")
	flag.Parse()

	// Connect to the database
	db, err := connectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to MySQL successfully!")

	// Handle migration commands
	if *migrateCmd != "" {
		if err := handleMigrations(db, *migrateCmd, *moduleName); err != nil {
			log.Fatalf("Migration error: %v", err)
		}
	}

	// Start the application (e.g., HTTP server)
	startApplication()
}

// connectToDatabase establishes a connection to the MySQL database
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}

// handleMigrations processes migration commands
func handleMigrations(db *sql.DB, migrateCmd, moduleName string) error {
	switch migrateCmd {
	case "up":
		return applyMigrations(db, moduleName)
	case "down":
		return rollbackMigrations(db, moduleName)
	default:
		return fmt.Errorf("invalid migration command: %s", migrateCmd)
	}
}

// applyMigrations applies migrations for a specific module or all modules
func applyMigrations(db *sql.DB, moduleName string) error {
	if moduleName != "" {
		// Run migrations for a specific module
		if err := migrations.RunForModule(db, moduleName, "up"); err != nil {
			return fmt.Errorf("failed to apply migrations for module %s: %v", moduleName, err)
		}
	} else {
		// Run migrations for all modules
		if err := migrations.RunAll(db); err != nil {
			return fmt.Errorf("failed to apply migrations (up): %v", err)
		}
	}
	fmt.Println("Migrations applied successfully!")
	return nil
}

// rollbackMigrations rolls back migrations for a specific module or all modules
func rollbackMigrations(db *sql.DB, moduleName string) error {
	if moduleName != "" {
		// Rollback migrations for a specific module
		if err := migrations.RunForModule(db, moduleName, "down"); err != nil {
			return fmt.Errorf("failed to rollback migrations for module %s: %v", moduleName, err)
		}
	} else {
		// Rollback migrations for all modules
		if err := migrations.RollbackAll(db); err != nil {
			return fmt.Errorf("failed to rollback migrations (down): %v", err)
		}
	}
	fmt.Println("Migrations rolled back successfully!")
	return nil
}

// startApplication starts the application (e.g., HTTP server)
func startApplication() {
	fmt.Println("Starting the application...")

	// Create a new ServeMux for routing
	router := http.NewServeMux()

	// Setup routes for the users module
	routes.SetupUsersRoutes(router)

	// Default route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to AutoVerse!")
	})

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
