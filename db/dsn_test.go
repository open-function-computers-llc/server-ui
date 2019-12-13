package db

import "testing"

func TestMariaDSNGeneration(t *testing.T) {
	_, err := GenerateMariaDSN("", "", "", "", "")
	if err == nil {
		t.Errorf("Expected error, as all arugments are required")
	}

	dsn, err := GenerateMariaDSN("dbuser", "dbpass", "dbname", "host", "3306")
	expected := "dbuser:dbpass@tcp(host:3306)/dbname?parseTime=true"
	if err != nil {
		t.Errorf("Expected a valid DSN to be generated")
	}
	if dsn != expected {
		t.Errorf("Expected DSN was not generated correctly.\nExpected: %s\nActual:   %s", expected, dsn)
	}
}
