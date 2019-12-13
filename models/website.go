package models

import "github.com/jmoiron/sqlx"

// Website a representation of a site in the database
type Website struct {
	ID  int    `db:"id"`
	URL string `db:"url"`
}

func (w Website) GetSchema(driver string) string {
	return `CREATE TABLE IF NOT EXISTS websites (
		id INTEGER PRIMARY KEY,
		url VARCHAR(255) DEFAULT ''
	);`
}

// CreateSite - store a new site instance in the DB
func CreateSite(url string, db *sqlx.DB) error {
	statement := "INSERT INTO websites (url) VALUES (?)"
	_, err := db.Exec(statement, url)
	if err != nil {
		return err
	}
	return nil
}