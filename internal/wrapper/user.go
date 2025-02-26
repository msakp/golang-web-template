package wrapper

import (
	"solution/internal/database/storage"
	"solution/internal/domain/dto"
	"solution/pkg/utils"
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
