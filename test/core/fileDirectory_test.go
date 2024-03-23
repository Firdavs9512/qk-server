package core_test

import (
	"os"
	"testing"

	"github.com/Firdavs9512/qk-server/config"
	"github.com/Firdavs9512/qk-server/core"
)

func TestConfigFileDirectory(t *testing.T) {
	// Create directory
	core.CreateDirectory(config.App.UploadUrl)

	// Check if directory exists
	if _, err := os.Stat(config.App.UploadUrl); os.IsNotExist(err) {
		t.Errorf("Directory was not created")
	}

	// Remove directory
	os.Remove(config.App.UploadUrl)
}

func TestCreateDirectory(t *testing.T) {
	// Create directory
	core.CreateDirectory("ExampleDirectory")

	// Check if directory exists
	if _, err := os.Stat("ExampleDirectory"); os.IsNotExist(err) {
		t.Errorf("Directory was not created")
	}

	// Remove directory
	os.Remove("ExampleDirectory")
}
