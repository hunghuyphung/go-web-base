package repository

import "gorm.io/gorm"

type RoleRepository interface {
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: db,
	}
}
