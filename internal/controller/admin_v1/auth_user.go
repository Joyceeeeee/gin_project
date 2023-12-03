package admin_v1

import (
	"github.com/gin-gonic/gin"
	"youdangzhe/internal/controller"
	"youdangzhe/internal/service/admin_auth"
	"youdangzhe/internal/validator"
	"youdangzhe/internal/validator/form"
)

type AdminUserController struct {
	controller.Api
}

func NewAdminUserController() *AdminUserController {
	return &AdminUserController{}
}

func (api AdminUserController) GetUserInfo(c *gin.Context) {
	result, err := admin_auth.NewAdminUserService().GetUserInfo(c.GetUint("user_id"))
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, result)
	return
}

func (api AdminUserController) Add(c *gin.Context) {
	// 初始化参数结构体
	IDForm := form.NewIDForm()
	//// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &IDForm); err != nil {
		return
	}
	result, err := admin_auth.NewAdminUserService().GetUserInfo(IDForm.ID)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, result)
	return
}

func (api AdminUserController) RegisterUser(c *gin.Context) {
	// 初始化参数结构体
	registerUserForm := form.NewRegisterForm()
	//// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &registerUserForm); err != nil {
		return
	}
	result, err := admin_auth.NewAdminUserService().AddUser(registerUserForm)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, result)
	return
}


