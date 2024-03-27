package http

import (
	"fmt"

	"github.com/Firdavs9512/qk-server/app/http/controllers"
	"github.com/Firdavs9512/qk-server/app/http/middleware"
	"github.com/Firdavs9512/qk-server/config"
	"github.com/kataras/iris/v12"
)

type Server struct{}

var Application *iris.Application

func (s *Server) Start() {
	// Create a new Iris application
	Application = iris.Default()

	// Configure
	Application.Use(iris.LimitRequestBodySize(config.App.MaxFileSize))
	Application.Use(middleware.UserAuthMiddleware())

	Application.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Ok!"})
	})

	// Register the controllers
	upload := Application.Party("/upload")
	{
		upload.Post("/file", func(ctx iris.Context) {
			new(controllers.FileUploadController).Post(ctx)
		})
		upload.Post("/files", func(ctx iris.Context) {
			new(controllers.FilesUploadController).Post(ctx)
		})
	}

	Application.Listen(fmt.Sprintf("%s:%d", config.App.AppHost, config.App.AppPort))
}

func RestartServer() {
	if Application != nil {
		Application.ConfigureHost(func(h *iris.Supervisor) {
			// Restart the server
			h.Server.Addr = fmt.Sprintf("%s:%d", config.App.AppHost, config.App.AppPort)
		})
	}
}
