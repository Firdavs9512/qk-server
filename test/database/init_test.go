package database_test

import (
	"os"
	"testing"

	"github.com/Firdavs9512/qk-server/config"
)

func TestDatabaseInit(t *testing.T) {
	// Test your database connection here
	config.Database.Init()

	// Test if the database is connected
	if config.Database.DB == nil {
		t.Errorf("Database is not connected")
	}

	// Delete database file if it exists
	if config.Database.GetConnection() == "sqlite" {
		os.Remove(config.Database.GetHost())
	}
}
