package http

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

type Server struct {
	Host string
	Port int
}

func (s *Server) Start() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello, World!"})
	})

	app.Listen(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
