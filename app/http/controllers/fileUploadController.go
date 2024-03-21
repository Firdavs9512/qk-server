package controllers

import "github.com/kataras/iris/v12"

type FileUploadController struct {
}

func (f *FileUploadController) Post(ctx iris.Context) int {
	message := ctx.PostValue("message")

	ctx.JSON(iris.Map{"message": message})

	return iris.StatusCreated
}
