package models

import "testing"

func TestCanCreateWebsite(t *testing.T) {
	db, _ := setUpDB("sqlite3")
	err := CreateSite("www.testing.com", db)
	if err != nil {
		t.Errorf("We should have been able to save a site to the DB. Error: " + err.Error())
	}
}

func TestCanFindWebsiteByURL(t *testing.T) {
	db, _ := setUpDB("sqlite3")
	CreateSite("www.testing.com", db)

	w := Website{}
	err := w.FindByURL("www.testing.com", db)
	if err != nil {
		t.Errorf("We expected to find this website in the database")
	}
	if w.ID == 0 {
		t.Errorf("The ID should be populated because it was created in the DB and retrieved")
	}

	nonExistingWebsite := Website{}
	nonExistingWebsite.FindByURL("www.idontexsitinthedatabase.com", db)
	if nonExistingWebsite.ID != 0 {
		t.Errorf("This should be zero")
	}
}
