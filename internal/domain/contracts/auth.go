package contracts

import "github.com/golang-jwt/jwt/v5"

type AuthService interface {
	GenerateToken(subject string) (string, error)
	GetSubject(tokenString *jwt.Token) (string, error)
	TokenIsFresh(tokenString *jwt.Token) (bool, error)
}
