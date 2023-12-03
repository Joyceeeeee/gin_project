package form

type LoginAuth struct {
	Email string `form:"email" json:"email"  binding:"required,min=5"` //  验证规则：必填，最小长度为5
	PassWord string `form:"password" json:"password"  binding:"required,min=6"` //  验证规则：必填，最小长度为6
}

func NewLoginForm() *LoginAuth {
	return &LoginAuth{}
}
