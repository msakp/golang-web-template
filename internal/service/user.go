package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/msakp/golang-web-template/internal/common/utils"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/infrastructure/wrapper"
)

var _ contracts.UserService = (*userService)(nil)

type userService struct {
	userRepo  contracts.UserRepository
	secretKey string
}

func NewUserService(ur contracts.UserRepository, secretKey string) *userService {
	return &userService{
		userRepo:  ur,
		secretKey: secretKey,
	}
}

func (s *userService) Register(u *dto.UserRegister) (token string, id uuid.UUID, err error) {
	_, err = s.userRepo.GetByEmail(u.Email)
	if err == nil {
		return token, id, errors.New("email already registered")
	}
	u.PasswordHashed = utils.HashPassword(u.PasswordUnhashed)

	createParams := wrapper.WithUserRegister(u)
	id, err = s.userRepo.Create(createParams)
	if err != nil {
		return token, id, err
	}
	token, err = utils.GenerateToken(createParams.Email, s.secretKey)
	return token, id, err

}

func (s *userService) Login(uLogin *dto.UserLogin) (token string, id uuid.UUID, err error) {
	user, err := s.userRepo.GetByEmail(uLogin.Email)
	if err != nil {
		return token, id, errors.New("no user registered on this email")
	}
	ok := utils.CompareHashAndPassword(user.Password, uLogin.PasswordUnHashed)
	if !ok {
		return token, id, errors.New("password mismatch")
	}
	token, err = utils.GenerateToken(uLogin.Email, s.secretKey)
	return token, id, err

}

func (s *userService) GetProfile(email string) (*dto.UserView, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return wrapper.ToUserView(user), nil
}
