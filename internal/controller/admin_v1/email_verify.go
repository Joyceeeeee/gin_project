package admin_v1

import (
	"github.com/gin-gonic/gin"
	"youdangzhe/internal/controller"
	"youdangzhe/internal/validator"
	"youdangzhe/internal/validator/form"
)

type EmailVerificationController struct {
	controller.Api
}

func NewEmailVerificationController() *EmailVerificationController {
	return &EmailVerificationController{}
}

func (api *EmailVerificationController) SendCode(c *gin.Context) {
	// 初始化参数结构体
	emailForm := form.NewEmailCodeForm()
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &emailForm); err != nil {
		return
	}
	// TODO 发送验证码
	//result, err := admin_auth.NewLoginService().Login(email.Email, loginForm.PassWord)
	// 根据业务返回值判断业务成功 OR 失败
	//if err != nil {
	//	api.Err(c, err)
	//	return
	//}

	api.Success(c, "send verify code success")
	return
}
