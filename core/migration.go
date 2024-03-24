package core

import (
	"github.com/Firdavs9512/qk-server/app/models"
	"github.com/Firdavs9512/qk-server/config"
)

func Migrate() {
	config.Database.DB.AutoMigrate(
		&models.AuthToken{},
		&models.Settings{},
		&models.Files{},
	)
}
