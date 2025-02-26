package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/wrapper"
	"github.com/msakp/golang-web-template/pkg/logger"
)

var _ contracts.AuthService = (*authService)(nil)

type authService struct {
	secretKey string
}

func NewAuthService(secretKey string) *authService {
	return &authService{
		secretKey: secretKey,
	}
}

func (s *authService) GenerateToken(ctx context.Context, username string) (string, *dto.HttpErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(3 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("failed to generate token with ERR: %s", err.Error()))
		return "", wrapper.InternalServerErr("Failed to create subject")
	}
	return tokenString, nil
}

func (s *authService) GetSubject(ctx context.Context, token *jwt.Token) (string, *dto.HttpErr) {
	sub, err := token.Claims.GetSubject()
	if err != nil {
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("failed to read token with ERR: %s", err.Error()))
		return "", wrapper.InternalServerErr("Failed to retrieve auth subject")
	}
	return sub, nil

}
