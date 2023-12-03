package form

type RegisterUser struct {
	UserName string `form:"user_name" json:"username"  binding:"required"` //  验证规则：必填，最小长度为5
	PassWord string `form:"password" json:"password"  binding:"required,min=6"` //  验证规则：必填，最小长度为6
	Email    string `form:"email" json:"email"  binding:"required,email"`
	VerifyCode string `form:"verify_code" json:"verify_code" binding:"required"`
	Location string `form:"location" json:"location" binding:"required"`
	RegisterWay uint `form:"register_way" json:"register_way" binding:"required"`
}

func NewRegisterForm() *RegisterUser {
	return &RegisterUser{}
}

