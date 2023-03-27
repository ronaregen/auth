package models

import "time"

type UserInstance struct {
	ID           uint      `gorm:"primaryKey" json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	InstanceName string
}
