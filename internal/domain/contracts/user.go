package contracts

import (
	"context"

	"github.com/google/uuid"
	"solution/internal/domain/dto"
)

type UserRepository interface {
	Exists(ctx context.Context, email string) bool
	Create(ctx context.Context, u *dto.UserRegister) (uuid.UUID, error)
	GetByEmail(ctx context.Context, email string) (*dto.UserView, error)
	GetById(ctx context.Context, id uuid.UUID) (*dto.UserView, error)
}

type UserService interface {
	Register(ctx context.Context, u *dto.UserRegister) (*dto.UserAuthResponse, *dto.HttpErr)
	Login(ctx context.Context, uLogin *dto.UserLogin) (*dto.UserAuthResponse, *dto.HttpErr)
	GetProfile(ctx context.Context, email string) (*dto.UserView, *dto.HttpErr)
}
