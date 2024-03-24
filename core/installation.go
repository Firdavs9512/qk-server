package core

import (
	"fmt"

	"github.com/Firdavs9512/qk-server/config"
	"github.com/fatih/color"
)

func StartInitiation() {
	// Init database
	config.Database.Init()

	// Init file directory
	InitDirectory()

	// Console texts
	startText()
	authorText()
	versionText()
}

func startText() {
	fmt.Print(`
   ____    __ __    _____                          
  / __ \  / //_/   / ___/___  ______   _____  _____
 / / / / / ,<      \__ \/ _ \/ ___/ | / / _ \/ ___/
/ /_/ / / /| |    ___/ /  __/ /   | |/ /  __/ /    
\___\_\/_/ |_|   /____/\___/_/    |___/\___/_/    
`)
	fmt.Print("\n")
}

func authorText() {
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("Created by: %s\n", red("Firdavs"))
}

func versionText() {
	fmt.Printf("Version: %s\n", config.App.Version)
}
