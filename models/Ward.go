package models

import "time"

type Ward struct {
	ID           uint      `gorm:"primaryKey" json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	DistrictCode int
	District     District `gorm:"foreignKey:DistrictCode;references:Code"`
	Code         int      `gorm:"unique"`
	Name         string
}
