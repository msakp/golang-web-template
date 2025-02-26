package wrapper

import (
	"github.com/msakp/golang-web-template/internal/database/storage"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/pkg/utils"
)

func UserRegisterWithCreateParams(u *dto.UserRegister) *storage.CreateUserParams {
	u.PasswordHashed = utils.HashPassword(u.PasswordUnhashed)
	return &storage.CreateUserParams{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.PasswordHashed,
	}
}

func UserWithView(u *storage.User) *dto.UserView {
	return &dto.UserView{
		Id:             u.ID,
		Name:           u.Name,
		Email:          u.Email,
		PasswordHashed: u.Password,
	}
}
