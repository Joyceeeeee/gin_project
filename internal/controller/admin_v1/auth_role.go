package admin_v1

import "youdangzhe/internal/controller"

type RoleController struct {
	controller.Api
}

func NewRoleController() *RoleController {
	return &RoleController{}
}
