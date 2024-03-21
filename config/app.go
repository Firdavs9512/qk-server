package config

type AppType struct {
	Name        string
	Version     string
	Author      string
	UploadUrl   string
	AppHost     string
	AppPort     int
	MaxFileSize int64
}

var (
	App = AppType{
		Name:        "QK Server",
		Version:     "0.0.1",
		Author:      "Firdavs",
		UploadUrl:   "uploads/",
		AppHost:     "localhost",
		AppPort:     2595,
		MaxFileSize: 10 << 20, // 10MB
	}
)
