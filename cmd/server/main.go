package main

import (
	"github.com/Firdavs9512/qk-server/app/http"
	"github.com/Firdavs9512/qk-server/config"
)

func main() {
	server := http.Server{
		Host: config.App.AppHost,
		Port: config.App.AppPort,
	}

	server.Start()
}
