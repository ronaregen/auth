package models

import "time"

type Province struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Code      int       `gorm:"unique"`
	Name      string
}
