package service

import (
	"context"
	"fmt"

	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/wrapper"
	"github.com/msakp/golang-web-template/pkg/logger"
	"github.com/msakp/golang-web-template/pkg/utils"
)

var _ contracts.UserService = (*userService)(nil)

type userService struct {
	userRepo    contracts.UserRepository
	authService contracts.AuthService
}

func NewUserService(ur contracts.UserRepository, as contracts.AuthService) *userService {
	return &userService{
		userRepo:    ur,
		authService: as,
	}
}

func (s *userService) Register(ctx context.Context, u *dto.UserRegister) (*dto.UserAuthResponse, *dto.HttpErr) {
	if s.userRepo.Exists(ctx, u.Email) {
		return nil, wrapper.NotFoundErr(dto.MsgUserAlreadyExists)
	}

	id, err := s.userRepo.Create(ctx, u)
	if err != nil {
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("failed to create user with ERR: %s", err.Error()))
		return nil, wrapper.InternalServerErr(err.Error())
	}

	token, httpErr := s.authService.GenerateToken(ctx, u.Email)
	if httpErr != nil {
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("generate token failed with ERR: %s", httpErr.Error()))
		return nil, httpErr
	}

	return &dto.UserAuthResponse{Id: id, Token: token}, nil

}

func (s *userService) Login(ctx context.Context, uLogin *dto.UserLogin) (*dto.UserAuthResponse, *dto.HttpErr) {
	user, err := s.userRepo.GetByEmail(ctx, uLogin.Email)
	if err != nil {
		return nil, wrapper.NotFoundErr(dto.MsgUserNotFound)
	}

	ok := utils.CompareHashAndPassword(user.PasswordHashed, uLogin.PasswordUnHashed)
	if !ok {
		return nil, wrapper.BadRequestErr(dto.MsgInvalidPassword)
	}

	token, httpErr := s.authService.GenerateToken(ctx, uLogin.Email)
	if httpErr != nil {
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("generate token failed with ERR: %s", httpErr.Error()))
		return nil, httpErr
	}

	return &dto.UserAuthResponse{Id: user.Id, Token: token}, nil

}

func (s *userService) GetProfile(ctx context.Context, email string) (*dto.UserView, *dto.HttpErr) {
	if !s.userRepo.Exists(ctx, email) {
		return nil, wrapper.NotFoundErr(dto.MsgUserNotFound)
	}
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("failed to get user from db with ERR: %s", err.Error()))
		return nil, wrapper.InternalServerErr(err.Error())
	}
	return user, nil
}
