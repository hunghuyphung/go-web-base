package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string
	Desc        string
	Level       int
	Status      string
	users       []User
	authorities []Authority `gorm:"many2many:role_authorities;"`
}
