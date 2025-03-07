package main

import (
	"auto_verse/migrations"
	"fmt"
	"os"
)

func main() {
	// Get the module name and migration description from the command line
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run cmd/create_migration/main.go <module_name> <migration_description>")
		return
	}
	moduleName := os.Args[1]
	migrationDescription := os.Args[2]

	// Call the CreateMigration function from the migrations package
	migrations.CreateMigration(moduleName, migrationDescription)
}
