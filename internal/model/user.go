package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name,omitempty"`
	Username string `gorm:"not null; unique" json:"username,omitempty"`
	Password string `gorm:"not null" json:"-"`
}
