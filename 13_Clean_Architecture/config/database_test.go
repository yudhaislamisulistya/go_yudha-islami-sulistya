package config

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	db, err := ConnectDB()

	// Assertion
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}
	if db == nil {
		t.Error("Database connection is nil")
	}
}

func TestMigrateDB(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform database migration
	err = MigrateDB(db)

	// Assertion
	if err != nil {
		t.Errorf("Failed to migrate database: %v", err)
	}
}
