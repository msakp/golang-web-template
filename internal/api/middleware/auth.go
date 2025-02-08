package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
)

func Auth(authService contracts.AuthService) fiber.Handler{

return func(c *fiber.Ctx) error{
	bearer := c.Get("Authorization")
	splited := strings.Split(bearer, " ")
	if len(splited) != 2{
		return c.Status(401).JSON(dto.HttpErr{Message: "Request Unauthorized"})
	}
	tokenString := splited[1]
	ok, err := authService.TokenIsFresh(tokenString)
	if err != nil {
		return c.Status(401).JSON(dto.HttpErr{Message: err.Error()})
	
	}
	if !ok {
		return c.Status(401).JSON(dto.HttpErr{Message: "token is expired"})
	}
	_, err = authService.GetSubFromToken(tokenString)
	if err != nil {
		return c.Status(401).JSON(dto.HttpErr{Message: err.Error()})
	}
	return c.Next()
}

}
