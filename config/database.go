package config

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	//"log"
)

// NewMySQLStorage creates a new MySQL connection
func SQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
		//log.Fatal(err)
	}

	return db, nil
}




