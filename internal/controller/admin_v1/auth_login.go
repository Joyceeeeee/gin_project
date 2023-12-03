package admin_v1

import (
	"github.com/gin-gonic/gin"
	"youdangzhe/internal/controller"
	"youdangzhe/internal/service/admin_auth"
	"youdangzhe/internal/validator"
	"youdangzhe/internal/validator/form"
)

type LoginController struct {
	controller.Api
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

// Login 用户登录
func (api LoginController) Login(c *gin.Context) {
	// 初始化参数结构体
	loginForm := form.NewLoginForm()
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &loginForm); err != nil {
		return
	}
	// 实际业务调用
	result, err := admin_auth.NewLoginService().Login(loginForm.Email, loginForm.PassWord)
	// 根据业务返回值判断业务成功 OR 失败
	if err != nil {
		api.Err(c, err)
		return
	}

	api.Success(c, result)
	return
}
