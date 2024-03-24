package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseType struct {
	DB         *gorm.DB
	connection string
	host       string
	// port       int
	// username   string
	// password   string
}

var (
	Database = DatabaseType{
		connection: "sqlite",
		host:       "database.sqlite",
	}
)

func (d *DatabaseType) Set(db *DatabaseType) {
	Database = *db
}

func (d *DatabaseType) GetConnection() string {
	return d.connection
}

func (d *DatabaseType) GetHost() string {
	return d.host
}

func (d *DatabaseType) Init() {
	var err error
	d.DB, err = gorm.Open(sqlite.Open(d.host), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connected!")
}
