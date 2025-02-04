package contracts

import (
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/infrastructure/database/sqlc/storage"
)

type UserRepository interface {
	Create(u *storage.CreateUserParams) error
	Get(email string) (*storage.User, error)
}

type UserService interface {
	Create(u *dto.UserRegister) error
	Get(email string) (*dto.UserView, error)
}
