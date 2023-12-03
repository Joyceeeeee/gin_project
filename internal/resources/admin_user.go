package resources

import (
	"github.com/jinzhu/copier"
)

type RegisterUserResources struct {
	ID       uint   `json:"id"`
	RegisterNum uint `json:"register_num"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewRegisterUserResources(data any) *RegisterUserResources {
	var adminUser RegisterUserResources
	err := copier.Copy(&adminUser, data)
	if err != nil {
		return nil
	}
	return &adminUser
}
