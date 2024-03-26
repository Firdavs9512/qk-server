package models

import "gorm.io/gorm"

type AuthToken struct {
	gorm.Model
	Name  string
	Token string
}
