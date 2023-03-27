package models

import "time"

type City struct {
	ID           uint      `gorm:"primaryKey" json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	ProvinceCode int
	Province     Province `gorm:"foreignKey:ProvinceCode;references:Code"`
	Code         int      `gorm:"unique"`
	Name         string
}
