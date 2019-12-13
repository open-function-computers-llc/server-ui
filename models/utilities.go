package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // db connections via sqlite
)

// GetSchemas - return all the SQL commands needed to migrate a database
func GetSchemas(driver string) string {
	var output string

	// do this for each model
	w := Website{}
	output = output + w.GetSchema(driver)

	return output
}

// used in tests to bootstrap an in memory sqlite db
func setUpDB(driver string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, ":memory:")
	if err != nil {
		return nil, err
	}
	db.MustExec(GetSchemas(driver))

	return db, nil
}
