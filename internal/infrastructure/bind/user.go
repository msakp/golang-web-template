package bind

import (
	"github.com/msakp/golang-web-template/domain/dto"
	"github.com/msakp/golang-web-template/internal/infrastructure/database/sqlc/storage"
)




func WithUserRegister(u *dto.UserRegister) storage.CreateUserParams{
	return storage.CreateUserParams{
		Name: u.Name,
		Email: u.Email,
		Password: u.PasswordUnhashed,
	}
}

