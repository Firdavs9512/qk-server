package core

import (
	"fmt"
	"os"

	"github.com/Firdavs9512/qk-server/config"
)

func InitDirectory() {
	CreateDirectory(config.App.UploadUrl)
}

// Create new directory
func CreateDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create directory if it doesn't exist
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	} else if err != nil {
		// If there was an error other than non-existence
		fmt.Println("Error checking directory:", err)
		return
	}
}
