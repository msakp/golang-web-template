package v1

import (
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
	u.Post("/sign-up", uh.Register)
	u.Post("/sign-in", uh.Login)

}

func (uh *userHandler) Create(c *fiber.Ctx) error {
	userRegister := new(dto.UserRegister)
	if err := c.BodyParser(userRegister); err != nil {
		return c.Status(400).JSON(fiber.Map{"err": err.Error()})
	}
	_, err := uh.userService.Register(userRegister)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"err": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "ok"})
}

// Register godoc
//
//	@Tags		users
//	@Summary	register new user
//	@Param		RequestBody	body	dto.UserRegister	true	"Registers new user and return access token"
//	@Accept		json
//	@Produce	json
//	@Success	201	{object}	dto.UserRegisterResponse
//	@Router		/user/sign-in [post]
func (uh *userHandler) Register(c *fiber.Ctx) error {
	userRegister := new(dto.UserRegister)
	if err := c.BodyParser(userRegister); err != nil {
		return c.Status(400).JSON(fiber.Map{"err": err.Error()})
	}
	token, err := uh.userService.Register(userRegister)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"err": err.Error()})
	}
	return c.Status(201).JSON(dto.UserRegisterResponse{Token: token})
}

func (uh *userHandler) Login(c *fiber.Ctx) error {
	return nil
}
