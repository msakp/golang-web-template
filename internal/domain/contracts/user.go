package contracts

import (
	"github.com/google/uuid"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/infrastructure/database/sqlc/storage"
)

type UserRepository interface {
	Create(u *storage.CreateUserParams) (uuid.UUID, error)
	GetByEmail(email string) (*storage.User, error)
	GetById(id uuid.UUID) (*storage.User, error)
}

type UserService interface {
	Register(u *dto.UserRegister) (token string, id uuid.UUID, err error)
	Login(uLogin *dto.UserLogin) (token string, id uuid.UUID, err error)
	GetProfile(email string) (*dto.UserView, error)
}
