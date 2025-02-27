package golangwebtemplate

import (
	"context"
	"solution/internal/wrapper"

	"github.com/gofiber/fiber/v2"
)

type someHandler struct{}


func (sh *someHandler) Dosmth(c *fiber.Ctx) error{
	var Data any
	if err := c.BodyParser(&Data); err != nil{
		httpErr := wrapper.BadRequestErr(err.Error())
		return c.Status(httpErr.HttpCode).JSON(httpErr)
	}
	if err := sh.validatorService.ValidateRequestData(&Data); err != nil{
		return c.Status(err.HttpCode).JSON(err)
	}
	}
}
