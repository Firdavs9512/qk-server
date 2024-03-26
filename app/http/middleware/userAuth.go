package middleware

import (
	"github.com/Firdavs9512/qk-server/app/models"
	"github.com/Firdavs9512/qk-server/config"
	"github.com/kataras/iris/v12"
)

type RequestHeader struct {
	Authorization string `header:"Authorization,required"`
}

func UserAuthMiddleware() iris.Handler {
	return func(ctx iris.Context) {
		var requestHeader RequestHeader
		if err := ctx.ReadHeaders(&requestHeader); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": "Invalid request"})
			return
		}

		if requestHeader.Authorization == "" {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{"message": "Unauthorized"})
			return
		}

		var token models.AuthToken
		if err := config.Database.DB.Where("token = ?", requestHeader.Authorization).First(&token).Error; err != nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{"message": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}
