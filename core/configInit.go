package core

import (
	"strconv"

	"github.com/Firdavs9512/qk-server/app/http"
	"github.com/Firdavs9512/qk-server/app/models"
	"github.com/Firdavs9512/qk-server/config"
)

// Local config files initialization in database
func ConfigInit() {
	var appConfig config.AppType
	appConfig.Name = config.App.Name
	appConfig.Author = config.App.Author
	// TODO: Change this to database
	appConfig.MaxFileSize = config.App.MaxFileSize

	// Application host
	var host *models.Settings
	config.Database.DB.Where("key = ?", "app_host").First(&host)
	if host == nil {
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
	config.Database.DB.Where("key = ?", "app_port").First(&port)
	if port == nil {
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
	config.Database.DB.Where("key = ?", "app_version").First(&version)
	if version == nil {
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
	config.Database.DB.Where("key = ?", "upload_url").First(&uploadUrl)
	if uploadUrl == nil {
		config.Database.DB.Create(&models.Settings{
			Key:   "upload_url",
			Value: config.App.UploadUrl,
		})
		appConfig.UploadUrl = config.App.UploadUrl
	} else {
		appConfig.UploadUrl = uploadUrl.Value
	}

	// Set the new config
	config.App = appConfig

	// Restart http server
	http.RestartServer()
}
