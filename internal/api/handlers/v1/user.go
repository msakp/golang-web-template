package v1

import (
	"time"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
)

type userHandler struct {
	userService contracts.UserService
}

func NewUserHandler(us contracts.UserService) *userHandler {
	return &userHandler{
		userService: us,
	}
}

func (uh *userHandler) Setup(r fiber.Router) {
	u := r.Group("/user")
	u.Get("/:email", uh.Get)
	u.Post("/sign-up", uh.Create)

}

func (uh *userHandler) Get(c *fiber.Ctx) error {
	email := c.Params("email")
	userView, err := uh.userService.Get(email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"err": err.Error()})
	}
	return c.Status(200).JSON(userView)

}

func (uh *userHandler) Create(c *fiber.Ctx) error {
	start := time.Now()
	userRegister := new(dto.UserRegister)
	if err := c.BodyParser(userRegister); err != nil {
		return c.Status(400).JSON(fiber.Map{"err": err.Error()})
	}
	err := uh.userService.Create(userRegister)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"err": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "ok", "time":  fmt.Sprintf("%s", time.Since(start))})
}
