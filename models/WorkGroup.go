package models

import "time"

type WorkGroup struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	GroupName string
}
