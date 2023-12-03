package form

type EmailVerifyVld struct {
	Email string `form:"email" json:"email"  binding:"required,email"`
}

func NewEmailCodeForm() *EmailVerifyVld {
	return &EmailVerifyVld{}
}
