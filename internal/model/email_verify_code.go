package model

type EmailVerifyCode struct {
	BaseModel
	Email    string `json:"email"`    // 邮箱
	VerifyCode string `json:"verify_code"`
}

func NewEmailVerificCode() *EmailVerifyCode {
	return &EmailVerifyCode{}
}

// TableName 获取表名
func (m *EmailVerifyCode) TableName() string {
	return "email_code"
}

func (m *EmailVerifyCode) CheckCodeVerified(code string) bool {
	return m.VerifyCode == code
}