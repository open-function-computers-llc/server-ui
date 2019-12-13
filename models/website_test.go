package models

import "testing"

func TestCanCreateWebsite(t *testing.T) {
	db, _ := setUpDB("sqlite3")
	err := CreateSite("www.testing.com", db)
	if err != nil {
		t.Errorf("We should have been able to save a site to the DB. Error: " + err.Error())
	}
}
