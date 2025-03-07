package users

import (
	"auto_verse/migrations"
	"database/sql"
)

func init() {
	// Register the users migration function
	migrations.Register(Migrate)
}

// Migrate runs all migrations for the users module
func Migrate(db *sql.DB) error {
	// Migration logic for the users module

	return nil
}
