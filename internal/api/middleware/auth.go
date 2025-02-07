package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/msakp/golang-web-template/internal/common/utils"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/spf13/viper"
)

func AuthMiddleware(secretKey string) fiber.Handler{

return func(c *fiber.Ctx) error{
	secretKey := viper.GetString("SECRET_KEY")
	bearer := c.Get("Authorization")
	splited := strings.Split(bearer, " ")
	if len(splited) != 2{
		return c.Status(401).JSON(dto.HttpErr{Message: "Request Unauthorized"})
	}
	tokenString := splited[1]
	ok, err := utils.TokenIsFresh(tokenString, secretKey)
	if err != nil {
		return c.Status(401).JSON(dto.HttpErr{Message: err.Error()})
	
	}
	if !ok {
		return c.Status(401).JSON(dto.HttpErr{Message: "token is expired"})
	}
	_, err = utils.GetSubFromToken(tokenString, secretKey)
	if err != nil {
		return c.Status(401).JSON(dto.HttpErr{Message: err.Error()})
	}
	return c.Next()
}

}
