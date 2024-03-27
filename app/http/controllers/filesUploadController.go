package controllers

import (
	"path/filepath"

	"github.com/Firdavs9512/qk-server/app/models"
	"github.com/Firdavs9512/qk-server/config"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type FilesUploadController struct{}

func (f *FilesUploadController) Post(ctx iris.Context) int {
	err := ctx.Request().ParseMultipartForm(ctx.Application().ConfigurationReadOnly().GetPostMaxMemory())
	if err != nil {
		ctx.JSON(iris.Map{"status": "error", "message": "Error parsing form data"})
		return iris.StatusInternalServerError
	}

	form := ctx.Request().MultipartForm
	files := form.File["files[]"]
	var fileNames []string

	for _, file := range files {
		uuidModel := uuid.New().String()
		filePath := filepath.Join(config.App.UploadUrl, uuidModel)

		_, err := ctx.SaveFormFile(file, filePath)
		if err != nil {
			ctx.JSON(iris.Map{"status": "error", "message": err.Error()})
			return iris.StatusInternalServerError
		} else {
			// Save file to database
			var files models.Files
			files.Uuid = uuidModel
			files.Name = file.Filename
			files.Path = filePath
			fileNames = append(fileNames, filePath)

			config.Database.DB.Create(&files)
		}
	}

	ctx.JSON(iris.Map{"status": "success", "message": "Files uploaded successfully", "files": fileNames})

	return iris.StatusCreated
}
