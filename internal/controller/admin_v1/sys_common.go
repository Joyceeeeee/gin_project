package admin_v1

import (
	"youdangzhe/internal/controller"
)

type CommonController struct {
	controller.Api
}

func NewCommonController() *CommonController {
	return &CommonController{}
}
