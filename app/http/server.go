package http

import (
	"fmt"

	"github.com/Firdavs9512/qk-server/app/http/controllers"
	"github.com/Firdavs9512/qk-server/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Server struct {
	Host string
	Port int
}

func (s *Server) Start() {
	// Create a new Iris application
	app := iris.Default()

	// Init database
	config.Database.Init()

	// Configure
	app.Use(iris.LimitRequestBodySize(config.App.MaxFileSize))

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello, World!"})
	})

	mvc := mvc.New(app.Party("/upload"))
	mvc.Handle(new(controllers.FileUploadController))

	app.Listen(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
