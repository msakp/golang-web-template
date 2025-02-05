package service

import (
	"errors"

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

func (s *userService) Register(u *dto.UserRegister) (token string, err error) {
	_, err = s.userRepo.Get(u.Email)
	if err == nil {
		return token, errors.New("email already registered")
	}
	createParams := wrapper.WithUserRegister(u)
	err = s.userRepo.Create(createParams)
	if err != nil{
		return token, err
	}
	token, err = utils.GenerateToken(createParams.Email, s.secretKey)
	return token, err

}

func (s *userService) Get(email string) (*dto.UserView, error) {
	user, err := s.userRepo.Get(email)
	if err != nil {
		return nil, err
	}
	return wrapper.ToUserView(user), nil
}
