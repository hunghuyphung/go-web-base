package controller

import "go-gin-web/internal/service"

type RoleController interface {
}

type RoleControllerImpl struct {
	roleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		roleService: roleService,
	}
}
