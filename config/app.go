package config

type AppType struct {
	Name      string
	Version   string
	Author    string
	UploadUrl string
}

var (
	App = AppType{
		Name:      "QK Server",
		Version:   "0.0.1",
		Author:    "Firdavs",
		UploadUrl: "uploads/",
	}
)
