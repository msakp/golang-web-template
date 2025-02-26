package contracts

import (
	"context"

	"github.com/google/uuid"
	"github.com/msakp/golang-web-template/internal/database/storage"
	"github.com/msakp/golang-web-template/internal/domain/dto"
)

type UserRepository interface {
	Create(ctx context.Context, u *storage.CreateUserParams) (uuid.UUID, error)
	GetByEmail(ctx context.Context, email string) (*storage.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*storage.User, error)
}

type UserService interface {
	Register(ctx context.Context, u *dto.UserRegister) (token string, id uuid.UUID, err error)
	Login(ctx context.Context, uLogin *dto.UserLogin) (token string, id uuid.UUID, err error)
	GetProfile(ctx context.Context, email string) (*dto.UserView, error)
}
