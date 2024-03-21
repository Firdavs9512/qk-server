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
	failures := 0
	for _, file := range files {
		_, err = ctx.SaveFormFile(file, filepath.Join(config.App.UploadUrl, file.Filename))
		if err != nil {
			failures++
			ctx.Writef("Error: %s\n", err.Error())
		}
	}
	ctx.Writef("%d files uploaded", len(files)-failures)

	return iris.StatusCreated
}
