package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "github.com/msakp/golang-web-template/docs"
	v1 "github.com/msakp/golang-web-template/internal/api/handlers/v1"
	"github.com/msakp/golang-web-template/internal/api/middleware"
	"github.com/msakp/golang-web-template/internal/config"
	"github.com/msakp/golang-web-template/internal/repository"
	"github.com/msakp/golang-web-template/internal/service"
	"github.com/msakp/golang-web-template/pkg/connections/postgres"
)

type App struct {
	Config *config.Config
	Fiber  *fiber.App
	DB     *postgres.DB
}

func NewApp(ctx context.Context) *App {
	var app App
	app.Config = config.New()
	app.init(ctx)
	return &app
}

func (app *App) Start() {
	app.Fiber.Listen(app.Config.ServerAddr)
}

func (app *App) init(ctx context.Context) {
	app.connectDB(ctx)
	app.engineSetup(ctx)
	app.handlersSetup()
}

func (app *App) connectDB(ctx context.Context) {
	app.DB = postgres.New(ctx, app.Config)
	app.DB.Migrate(ctx)
}

func (app *App) engineSetup(ctx context.Context) {
	app.Fiber = fiber.New()
	app.Fiber.Use(recover.New())
	app.Fiber.Use(logger.New())
	app.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
	}))
	app.Fiber.Use(middleware.CustomContext(ctx))
}

func (app *App) CloseConnections(ctx context.Context) {
	app.DB.Close(ctx)
}

func (app *App) handlersSetup() {
	// route groups
	apiV1 := app.Fiber.Group("/api/v1")

	//api global routes
	apiV1.Get("docs/*", swagger.HandlerDefault)

	// global repos

	// global services
	authService := service.NewAuthService(app.Config.SecretKey)
	validatorService := service.NewValidatorService()

	// repos
	userRepo := repository.NewUserRepository(app.DB)

	// services
	userService := service.NewUserService(userRepo, authService)

	//handlers
	userHandler := v1.NewUserHandler(userService, authService, validatorService)

	//handlers setup
	userHandler.Setup(apiV1, app.Config.SecretKey)

}
