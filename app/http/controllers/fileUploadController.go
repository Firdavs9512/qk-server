package controllers

import (
	"path/filepath"

	"github.com/Firdavs9512/qk-server/config"
	"github.com/kataras/iris/v12"
)

type FileUploadController struct {
}

func (f *FileUploadController) Post(ctx iris.Context) int {
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return iris.StatusInternalServerError
	}

	form := ctx.Request().MultipartForm

	files := form.File["files[]"]
	for _, file := range files {
		_, err = ctx.SaveFormFile(file, filepath.Join(config.App.UploadUrl, file.Filename))
		if err != nil {
			ctx.JSON(iris.Map{"status": "error", "message": err.Error()})
			return iris.StatusInternalServerError
		}
	}

	ctx.JSON(iris.Map{"status": "success", "message": "Files uploaded successfully"})

	return iris.StatusCreated
}
