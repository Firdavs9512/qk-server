package core

import (
	"fmt"
	"strconv"

	"github.com/Firdavs9512/qk-server/app/http"
	"github.com/Firdavs9512/qk-server/app/models"
	"github.com/Firdavs9512/qk-server/config"
	"github.com/Firdavs9512/qk-server/utils"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

// Local config files initialization in database
func ConfigInit() {
	var appConfig config.AppType
	appConfig.Name = config.App.Name
	appConfig.Author = config.App.Author

	// Max file size
	var maxFileSize *models.Settings
	mResult := config.Database.DB.Where("key = ?", "max_file_size").First(&maxFileSize)
	if mResult.Error == gorm.ErrRecordNotFound {
		config.Database.DB.Create(&models.Settings{
			Key:   "max_file_size",
			Value: strconv.FormatInt(config.App.MaxFileSize, 10),
		})
		appConfig.MaxFileSize = config.App.MaxFileSize
	} else {
		number, err := strconv.ParseInt(maxFileSize.Value, 10, 64)
		if err != nil {
			number = int64(config.App.MaxFileSize)
		}
		appConfig.MaxFileSize = number
	}

	// Application host
	var host *models.Settings
	hResult := config.Database.DB.Where("key = ?", "app_host").First(&host)
	if hResult.Error == gorm.ErrRecordNotFound {
		config.Database.DB.Create(&models.Settings{
			Key:   "app_host",
			Value: config.App.AppHost,
		})
		appConfig.AppHost = config.App.AppHost
	} else {
		appConfig.AppHost = host.Value
	}

	// Application port
	var port *models.Settings
	pResult := config.Database.DB.Where("key = ?", "app_port").First(&port)
	if pResult.Error == gorm.ErrRecordNotFound {
		config.Database.DB.Create(&models.Settings{
			Key:   "app_port",
			Value: strconv.Itoa(config.App.AppPort),
		})
		appConfig.AppPort = config.App.AppPort
	} else {
		number, err := strconv.Atoi(port.Value)
		if err != nil {
			number = config.App.AppPort
		}
		appConfig.AppPort = number
	}

	// Application version
	var version *models.Settings
	vResult := config.Database.DB.Where("key = ?", "app_version").First(&version)
	if vResult.Error == gorm.ErrRecordNotFound {
		config.Database.DB.Create(&models.Settings{
			Key:   "app_version",
			Value: config.App.Version,
		})
		appConfig.Version = config.App.Version
	} else {
		appConfig.Version = version.Value
	}

	// Application Upload URL
	var uploadUrl *models.Settings
	uResult := config.Database.DB.Where("key = ?", "upload_url").First(&uploadUrl)
	if uResult.Error == gorm.ErrRecordNotFound {
		config.Database.DB.Create(&models.Settings{
			Key:   "upload_url",
			Value: config.App.UploadUrl,
		})
		appConfig.UploadUrl = config.App.UploadUrl
	} else {
		appConfig.UploadUrl = uploadUrl.Value
	}

	// Check auth token if not exists create one
	var count int64
	config.Database.DB.Model(&models.AuthToken{}).Count(&count)
	if count == 0 {
		token := utils.RandomString(32)
		config.Database.DB.Create(&models.AuthToken{
			Token: token,
			Name:  "Default",
		})

		fmt.Printf("Default Auth Token: %s\n", color.HiYellowString(token))
	}

	// Set the new config
	config.App = appConfig

	// Restart http server
	http.RestartServer()
}
