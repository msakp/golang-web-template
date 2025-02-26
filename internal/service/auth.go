package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
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

func (s *authService) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(3 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.secretKey))
	return tokenString, err
}

func (s *authService) GetSubject(token *jwt.Token) (string, error) {
	return token.Claims.GetSubject()

}

func (s *authService) TokenIsFresh(token *jwt.Token) (bool, error) {
	expdate, err := token.Claims.GetExpirationTime()
	if err != nil {
		return false, err
	}
	now := time.Now()
	return expdate.After(now), nil
}
