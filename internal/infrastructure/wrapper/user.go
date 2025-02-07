package wrapper

import (
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/infrastructure/database/sqlc/storage"
)

func WithUserRegister(u *dto.UserRegister) *storage.CreateUserParams {
	return &storage.CreateUserParams{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.PasswordHashed,
	}
}

func ToUserView(u *storage.User) *dto.UserView {
	return &dto.UserView{
		Name:  u.Name,
		Email: u.Email,
		PasswordHashed: u.Password,
	}
}
