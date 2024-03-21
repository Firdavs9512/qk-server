package config

type AppType struct {
	Name      string
	Version   string
	Author    string
	UploadUrl string
	AppHost   string
	AppPort   int
}

var (
	App = AppType{
		Name:      "QK Server",
		Version:   "0.0.1",
		Author:    "Firdavs",
		UploadUrl: "uploads/",
		AppHost:   "localhost",
		AppPort:   2595,
	}
)
