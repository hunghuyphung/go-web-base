package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	FullName string
	Email    string
	RoleName string
	Level    int
	RoleId   int
}
