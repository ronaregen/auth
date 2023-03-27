package models

import "time"

type District struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	CityCode  int
	City      City `gorm:"foreignKey:CityCode;references:Code"`
	Code      int  `gorm:"unique"`
	Name      string
}
