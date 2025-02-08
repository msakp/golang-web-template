package v1

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/msakp/golang-web-template/internal/api/middleware"
	"github.com/msakp/golang-web-template/internal/common/utils"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
)

type userHandler struct {
	userService contracts.UserService
	secretKey string
}

func NewUserHandler(us contracts.UserService, secretKey string) *userHandler {
	return &userHandler{
		userService: us,
		secretKey: secretKey,
	}
}

func (uh *userHandler) Setup(r fiber.Router) {
	u := r.Group("/user")
	u.Post("/sign-up", uh.Register)
	u.Post("/sign-in", uh.Login)
	u.Get("/me", middleware.Auth(uh.secretKey), uh.GetInfo)

}


// RegisterUser godoc
//
//	@Tags		users
//	@Summary	register new user
//	@Param		RequestBody	body	dto.UserRegister	true	"Registers new user and returns access token"
//	@Accept		json
//	@Produce	json
//	@Success	201	{object}	dto.UserAuthResponse
//	@Failure	400	{object}	dto.HttpErr
//	@Router		/user/sign-up [post]
func (uh *userHandler) Register(c *fiber.Ctx) error {
	userRegister := new(dto.UserRegister)
	if err := c.BodyParser(userRegister); err != nil {
		return c.Status(400).JSON(dto.HttpErr{Message: err.Error()})
	}
	token, id, err := uh.userService.Register(c.Context(), userRegister)
	if err != nil {
		return c.Status(400).JSON(dto.HttpErr{Message: err.Error()})
	}
	return c.Status(201).JSON(dto.UserAuthResponse{Token: token, Id: id})
}


// LoginUser godoc
//
//	@Tags		users
//	@Summary	login existed user
//	@Param		RequestBody	body	dto.UserLogin	true	"Logins existed user and returns access token"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.UserAuthResponse
//	@Failure	400	{object}	dto.HttpErr
//	@Router		/user/sign-in [post]
func (uh *userHandler) Login(c *fiber.Ctx) error {
	userLogin := new(dto.UserLogin)
	if err := c.BodyParser(userLogin); err != nil {
		return c.Status(400).JSON(dto.HttpErr{Message: err.Error()})
	}
	token, id, err := uh.userService.Login(c.Context(), userLogin)
	if err != nil {
		return c.Status(400).JSON(dto.HttpErr{Message: err.Error()})
	}
	return c.Status(200).JSON(dto.UserAuthResponse{Token: token, Id: id})

}



// Profile godoc
//
//	@Tags		users
//	@Summary	get user profile data
//	@Security	Bearer
//  @Param Authorization header string true "access token 'Bearer {token}'"
//	@Produce	json
//	@Success	200	{object}	dto.UserView
//	@Failure	401	{object}	dto.HttpErr
//	@Router		/user/me [get]
func (uh *userHandler) GetInfo(c *fiber.Ctx) error{
	tokenString := strings.Split(c.Get("Authorization"), " ")
	email, _ := utils.GetSubFromToken(tokenString[1], uh.secretKey)
	user, err := uh.userService.GetProfile(c.Context(), email)
	if err != nil{
		return c.Status(400).JSON(dto.HttpErr{Message: err.Error()})
	}
	return c.Status(200).JSON(user)
}
