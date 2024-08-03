package service

import "go-gin-web/internal/repository"

type RoleService interface {
}

type RoleServiceImpl struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return &RoleServiceImpl{roleRepository: roleRepository}
}
