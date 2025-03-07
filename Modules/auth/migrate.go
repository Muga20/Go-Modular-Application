package auth

import (
	"database/sql"
	"auto_verse/migrations"
)

func init() {
	// Register the auth migration function
	migrations.Register(Migrate)
}

// Migrate runs all migrations for the auth module
func Migrate(db *sql.DB) error {
	// Migration logic for the auth module
	return nil
}
