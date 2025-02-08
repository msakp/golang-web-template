package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/msakp/golang-web-template/internal/common/utils"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/infrastructure/wrapper"
)

var _ contracts.UserService = (*userService)(nil)

type userService struct {
	authService contracts.AuthService
	userRepo  contracts.UserRepository
	secretKey string
}

func NewUserService(as contracts.AuthService, ur contracts.UserRepository, secretKey string) *userService {
	return &userService{
		authService: as,
		userRepo:  ur,
		secretKey: secretKey,
	}
}

func (s *userService) Register(ctx context.Context, u *dto.UserRegister) (token string, id uuid.UUID, err error) {
	_, err = s.userRepo.GetByEmail(ctx, u.Email)
	if err == nil {
		return token, id, errors.New("email already registered")
	}
	u.PasswordHashed = utils.HashPassword(u.PasswordUnhashed)

	createParams := wrapper.WithUserRegister(u)
	id, err = s.userRepo.Create(ctx, createParams)
	if err != nil {
		return token, id, err
	}
	token, err = s.authService.GenerateToken(createParams.Email)
	return token, id, err

}

func (s *userService) Login(ctx context.Context, uLogin *dto.UserLogin) (token string, id uuid.UUID, err error) {
	user, err := s.userRepo.GetByEmail(ctx, uLogin.Email)
	if err != nil {
		return token, id, errors.New("no user registered on this email")
	}
	ok := utils.CompareHashAndPassword(user.Password, uLogin.PasswordUnHashed)
	if !ok {
		return token, id, errors.New("password mismatch")
	}
	token, err = s.authService.GenerateToken(uLogin.Email)
	return token, id, err

}

func (s *userService) GetProfile(ctx context.Context, email string) (*dto.UserView, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return wrapper.ToUserView(user), nil
}
