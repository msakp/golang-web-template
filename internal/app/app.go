package app

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "github.com/msakp/golang-web-template/docs"
	v1 "github.com/msakp/golang-web-template/internal/api/handlers/v1"
	"github.com/msakp/golang-web-template/internal/api/middleware"
	"github.com/msakp/golang-web-template/internal/common/config"
	"github.com/msakp/golang-web-template/internal/infrastructure/database"
	"github.com/msakp/golang-web-template/internal/repository"
	"github.com/msakp/golang-web-template/internal/service"
)

type App struct {
	Config *config.Config
	Fiber  *fiber.App
	DB     *database.Pg
}

func NewApp(ctx context.Context) *App {
	var app App
	app.Config = config.New()
	app.init(ctx)
	return &app
}

func (app *App) Start() {
	log.Fatal(app.Fiber.Listen(app.Config.ServerAddr))
}

func (app *App) init(ctx context.Context) {
	app.connectDB(ctx)
	app.engineSetup(ctx)
	app.handlersSetup()
}

func (app *App) connectDB(ctx context.Context) {
	app.DB = database.NewPg(ctx, app.Config)
	app.DB.Migrate()
}

func (app *App) engineSetup(ctx context.Context) {
	app.Fiber = fiber.New()
	app.Fiber.Use(recover.New())
	app.Fiber.Use(logger.New())
	app.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "http://0.0.0.0:3000",
		AllowMethods: "*",
	}))
	app.Fiber.Use(middleware.CustomContext(ctx))
}

func (app *App) handlersSetup() {
	// route groups
	apiV1 := app.Fiber.Group("/api/v1")

	//add swagger spec
	apiV1.Get("docs/*", swagger.HandlerDefault)

	// user
	userRepo := repository.NewUserRepository(app.DB)
	authService := service.NewAuthService(app.Config.SecretKey)
	userService := service.NewUserService(authService, userRepo)
	userHandler := v1.NewUserHandler(userService, authService)
	userHandler.Setup(apiV1)

}
