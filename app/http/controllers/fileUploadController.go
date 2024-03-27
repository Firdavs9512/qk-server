package controllers

import (
	"path/filepath"

	"github.com/Firdavs9512/qk-server/app/models"
	"github.com/Firdavs9512/qk-server/config"
	"github.com/google/uuid"
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

	_, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(iris.Map{"status": "error", "message": err.Error()})
		return iris.StatusInternalServerError
	}

	uuidModel := uuid.New().String()
	filePath := filepath.Join(config.App.UploadUrl, uuidModel)

	// Save file to the server
	_, err = ctx.SaveFormFile(fileHeader, filePath)
	if err != nil {
		ctx.JSON(iris.Map{"status": "error", "message": err.Error()})
		return iris.StatusInternalServerError
	}

	var file models.Files
	file.Uuid = uuidModel
	file.Name = fileHeader.Filename
	file.Path = filePath

	config.Database.DB.Create(&file)

	ctx.JSON(iris.Map{"status": "success", "message": "File uploaded successfully", "file": filePath})
	return iris.StatusCreated
}
