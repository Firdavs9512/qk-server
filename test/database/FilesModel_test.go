package database_test

import (
	"os"
	"testing"

	"github.com/Firdavs9512/qk-server/config"
	"github.com/Firdavs9512/qk-server/core"
)

func TestFilesModel(t *testing.T) {
	config.Database.Init()

	// Migrate database models
	core.Migrate()

	// Files model exists
	if !config.Database.DB.Migrator().HasTable("files") {
		t.Errorf("Table files does not exist")
	}

	// Delete database file if it exists
	if config.Database.GetConnection() == "sqlite" {
		os.Remove(config.Database.GetHost())
	}
}
