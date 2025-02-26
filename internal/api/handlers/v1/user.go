package v1

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/wrapper"
)

type userHandler struct {
	userService      contracts.UserService
	authService      contracts.AuthService
	validatorService contracts.ValidatorService
}

func NewUserHandler(us contracts.UserService, as contracts.AuthService, vs contracts.ValidatorService) *userHandler {
	return &userHandler{
		userService:      us,
		authService:      as,
		validatorService: vs,
	}
}

func (uh *userHandler) Setup(r fiber.Router, secretKey string) {
	authMiddle := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secretKey)},
		ContextKey: "subject",
	})
	u := r.Group("/users")
	u.Post("/sign-up", uh.Register)
	u.Post("/sign-in", uh.Login)
	u.Get("/me", authMiddle, uh.GetInfo)

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
//	@Failure	500	{object}	dto.HttpErr
//	@Router		/users/sign-up [post]
func (uh *userHandler) Register(c *fiber.Ctx) error {
	userRegister := new(dto.UserRegister)
	if err := c.BodyParser(userRegister); err != nil {
		httpErr := wrapper.BadRequestErr(err.Error())
		return c.Status(httpErr.HttpCode).JSON(httpErr)
	}
	if err := uh.validatorService.ValidateRequestData(&userRegister); err != nil {
		return c.Status(err.HttpCode).JSON(err)
	}
	resp, err := uh.userService.Register(c.Context(), userRegister)
	if err != nil {
		return c.Status(err.HttpCode).JSON(err)
	}
	return c.Status(201).JSON(resp)
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
//	@Failure	404	{object}	dto.HttpErr
//	@Failure	500	{object}	dto.HttpErr
//	@Router		/user/sign-in [post]
func (uh *userHandler) Login(c *fiber.Ctx) error {
	userLogin := new(dto.UserLogin)
	if err := c.BodyParser(userLogin); err != nil {
		httpErr := wrapper.BadRequestErr(err.Error())
		return c.Status(httpErr.HttpCode).JSON(httpErr)
	}
	if err := uh.validatorService.ValidateRequestData(&userLogin); err != nil {
		return c.Status(err.HttpCode).JSON(err)
	}
	resp, err := uh.userService.Login(c.UserContext(), userLogin)
	if err != nil {
		return c.Status(err.HttpCode).JSON(err)
	}
	return c.Status(200).JSON(resp)

}

// Profile godoc
//
//	@Tags		users
//	@Summary	get user profile data
//	@Security	Bearer
//	@Param		Authorization	header	string	true	"access token 'Bearer {token}'"
//	@Produce	json
//	@Success	200	{object}	dto.UserView
//	@Failure	401	{object}	dto.HttpErr
//	@Failure	500	{object}	dto.HttpErr
//	@Router		/user/me [get]
func (uh *userHandler) GetInfo(c *fiber.Ctx) error {
	token := c.Locals("subject").(*jwt.Token)
	email, err := uh.authService.GetSubject(c.UserContext(), token)

	if err != nil {
		return c.Status(err.HttpCode).JSON(err)
	}
	user, err := uh.userService.GetProfile(c.UserContext(), email)
	if err != nil {
		return c.Status(err.HttpCode).JSON(err)
	}
	return c.Status(200).JSON(user)
}
