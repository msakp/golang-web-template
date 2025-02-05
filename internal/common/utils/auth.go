package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func ParseToken(tokenString, secretKey string) (jwt.Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("bad signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	return token.Claims, err
}

func GetSubFromToken(tokenString, secretKey string) (string, error) {
	token, err := ParseToken(tokenString, secretKey)
	if err != nil {
		return "", err
	}
	username, err := token.GetSubject()
	if err != nil {
		return "", err
	}
	return username, nil

}

func TokenIsFresh(tokenString, secretKey string) (bool, error) {
	token, err := ParseToken(tokenString, secretKey)
	if err != nil {
		return true, err
	}
	expdate, err := token.GetExpirationTime()
	now := time.Now()
	return expdate.After(now), nil
}
