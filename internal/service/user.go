package service

import (
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/infrastructure/wrapper"
)

var _ contracts.UserService = (*userService)(nil)

type userService struct {
	userRepo contracts.UserRepository
}

func NewUserService(ur contracts.UserRepository) *userService {
	return &userService{
		userRepo: ur,
	}
}

func (s *userService) Create(u *dto.UserRegister) error {
	createParams := wrapper.WithUserRegister(u)

	return s.userRepo.Create(createParams)
}

func (s *userService) Get(email string) (*dto.UserView, error) {
	user, err := s.userRepo.Get(email)
	if err != nil {
		return nil, err
	}
	return wrapper.ToUserView(user), nil
}
