package models

import "gorm.io/gorm"

type Files struct {
	gorm.Model
	Uuid string
	Name string
	Path string
}
