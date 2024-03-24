package main

import (
	"github.com/Firdavs9512/qk-server/app/http"
	"github.com/Firdavs9512/qk-server/core"
)

func main() {
	// Start installation
	core.StartInitiation()

	server := http.Server{}

	server.Start()
}
