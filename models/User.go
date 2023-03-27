package models

import (
	"time"
)

type User struct {
	ID             uint      `gorm:"primaryKey" json:"-"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	Username       string    `gorm:"unique"`
	Name           string
	Password       string `json:"-"`
	UserRoleId     int
	UserRole       UserRole `gorm:"foreignKey:UserRoleId"`
	UserInstanceId int
	UserInstance   UserInstance `gorm:"foreignKey:UserInstanceId"`
	WorkGroupId    int
	WorkGroup      WorkGroup
	ProvinceCode   int
	Province       Province `gorm:"foreignKey:ProvinceCode;references:Code"`
	CityCode       int
	City           City `gorm:"foreignKey:CityCode;references:Code"`
	DistrictCode   int
	District       District `gorm:"foreignKey:DistrictCode;references:Code"`
	WardCode       int
	Ward           Ward `gorm:"foreignKey:WardCode;references:Code"`
}
