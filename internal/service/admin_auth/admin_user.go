package admin_auth

import (
	"youdangzhe/internal/model"
	"youdangzhe/internal/pkg/errors"
	"youdangzhe/internal/resources"
	"youdangzhe/internal/service"
	"youdangzhe/internal/validator/form"
)

// AdminUserService 授权服务
type AdminUserService struct {
	service.Base
}

func NewAdminUserService() *AdminUserService {
	return &AdminUserService{}
}

// GetUserInfo 获取用户信息
func (a *AdminUserService) GetUserInfo(id uint) (*resources.RegisterUserResources, error) {
	// 查询用户是否存在
	//usersModel := model.NewUsers()
	//user := usersModel.GetUserById(id)
	user := model.User{
		IsAdmin:        0,
		Username:       "youdangzhe",
		Password:       "",
		Email:          "test@mail.com",
		EmailVerified:  true,
		Location:       "Dutch",
		Rank:           0,
		YouDangCash:    0,
		RegisterWay:    1,
		RegisterNum:    0,
		Description:    "",
	}
	return resources.NewRegisterUserResources(user), nil
	//if user != nil {
	//	return resources.NewRegisterUserResources(user), nil
	//}
	//return nil, errors.NewBusinessError(errors.FAILURE, "获取用户信息失败")
}

// AddUser 新增用户
func (a *AdminUserService) AddUser(registerUser *form.RegisterUser) (*resources.RegisterUserResources, error) {
	// 查询用户是否存在
	usersModel := model.UserFromRegisterForm(*registerUser)
	user, err := usersModel.Register()
	if err == nil {
		return resources.NewRegisterUserResources(user), nil
	}
	return nil, errors.NewBusinessError(errors.FAILURE, "注册用户失败")
}
