package model

import (
	"golang.org/x/crypto/bcrypt"
	"youdangzhe/internal/validator/form"
)

type User struct {
	BaseModel
	IsAdmin  int8   `json:"is_admin"` // 是否是总管理员
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Description string `json:"description"`
	Email    string `json:"email"`    // 邮箱
	EmailVerified bool   `json:"email_verified"`
	Location string `json:"locations"`
	Rank uint `json:"rank"` // 等级
	YouDangCash uint `json:"youdang_cash"`
	RegisterWay uint `json:"register_way"`
	RegisterNum uint `json:"register_num"`
}

func NewUsers() *User {
	return &User{}
}

func UserFromRegisterForm(registerUser form.RegisterUser) *User {
	user := User{
		IsAdmin:        0, // 设置默认值或其他逻辑
		Username:       registerUser.UserName,
		Password:       registerUser.PassWord,
		Email:          registerUser.Email,
		EmailVerified:  false, // 根据需要设置默认值
		Location:       registerUser.Location,
		Rank:           0, // 设置默认值或其他逻辑
		YouDangCash:    0, // 设置默认值或其他逻辑
		RegisterWay:    registerUser.RegisterWay,
		RegisterNum:    0, // 设置默认值或其他逻辑
		Description:    "", // 设置默认值或其他逻辑
	}
	return &user
}

// TableName 获取表名
func (m *User) TableName() string {
	return "users"
}

// GetUserById 根据uid获取用户信息
func (m *User) GetUserById(id uint) *User {
	if err := m.DB().First(m, id).Error; err != nil {
		return nil
	}
	return m
}

// Register 用户注册，写入到DB
func (m *User) Register() (*User, error) {
	m.Password, _ = m.PasswordHash(m.Password)
	//result := m.DB().Create(m)
	return m, nil
}

// PasswordHash 密码hash并自动加盐
func (m *User) PasswordHash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

// ComparePasswords 比对用户密码是否正确
func (m *User) ComparePasswords(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

// ChangePassword 修改密码
func (m *User) ChangePassword() error {
	m.Password, _ = m.PasswordHash(m.Password)
	return m.DB(m).Update("password", m.Password).Error
}

// GetUserInfoByEmail 根据名称获取用户信息
func (m *User) GetUserInfoByEmail(email string) *User {
	if err := m.DB().Where("email", email).First(m).Error; err != nil {
		return nil
	}
	return m
}
