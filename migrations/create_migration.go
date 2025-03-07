package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

// MigrationTemplate represents the data needed to generate a migration file
type MigrationTemplate struct {
	MigrationName string
	Timestamp     string
}

// CreateMigration creates a new migration file for the specified module
func CreateMigration(moduleName, migrationDescription string) {
	// Define the migration directory
	migrationDir := filepath.Join("Modules", moduleName, "migrations")

	// Create the migrations directory if it doesn't exist
	if err := os.MkdirAll(migrationDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create migrations directory: %v\n", err)
		return
	}

	// Generate a timestamp for the migration file name
	timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS format

	// Define the migration file names (up and down)
	upFileName := fmt.Sprintf("%s_%s.up.sql", timestamp, migrationDescription)
	downFileName := fmt.Sprintf("%s_%s.down.sql", timestamp, migrationDescription)

	upFilePath := filepath.Join(migrationDir, upFileName)
	downFilePath := filepath.Join(migrationDir, downFileName)

	// Define template data
	data := MigrationTemplate{
		MigrationName: migrationDescription,
		Timestamp:     timestamp,
	}

	// Create the up migration file from the template
	if err := createFileFromTemplate(upFilePath, upMigrationTemplate, data); err != nil {
		fmt.Printf("Failed to create up migration file: %v\n", err)
		return
	}

	// Create the down migration file from the template
	if err := createFileFromTemplate(downFilePath, downMigrationTemplate, data); err != nil {
		fmt.Printf("Failed to create down migration file: %v\n", err)
		return
	}

	fmt.Printf("Migration files created:\n- %s\n- %s\n", upFilePath, downFilePath)
}

// createFileFromTemplate creates a file from a template
func createFileFromTemplate(path, tmpl string, data MigrationTemplate) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	t := template.Must(template.New("").Parse(tmpl))
	return t.Execute(file, data)
}

// Template for up migration files
var upMigrationTemplate = `-- {{.Timestamp}}_{{.MigrationName}}.up.sql
-- Add your SQL statements here

-- Example: Create a table
CREATE TABLE {{.MigrationName}} (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
`

// Template for down migration files
var downMigrationTemplate = `-- {{.Timestamp}}_{{.MigrationName}}.down.sql
-- Add your SQL statements here

-- Example: Drop the table
DROP TABLE IF EXISTS {{.MigrationName}};
`
