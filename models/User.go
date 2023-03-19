package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Username   string
	Name       string
	Password   string `json:"-"`
}
