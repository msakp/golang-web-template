package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	v1 "github.com/msakp/golang-web-template/internal/api/handlers/v1"
	"github.com/msakp/golang-web-template/internal/config"
	"github.com/msakp/golang-web-template/internal/infrastructure/database"
	"github.com/msakp/golang-web-template/internal/repository"
	"github.com/msakp/golang-web-template/internal/service"
)

type App struct {
	Config *config.Config
	Fiber  *fiber.App
	DB     *database.Pg
}

func NewApp() *App {
	var app App
	app.Config = config.New()
	app.init()
	return &app
}

func (app *App) Start() {
	log.Fatal(app.Fiber.Listen(app.Config.ServerAddr))
}

func (app *App) init() {
	app.connectDB()
	app.engineSetup()
	app.handlersSetup()
}

func (app *App) connectDB() {
	app.DB = database.NewPg(app.Config)
}

func (app *App) engineSetup() {
	app.Fiber = fiber.New()
	app.Fiber.Use(recover.New())
	app.Fiber.Use(logger.New())
	app.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "http://0.0.0.0:3000",
		AllowMethods: "*",
	}))
}

func (app *App) handlersSetup() {
	// route groups
	apiV1 := app.Fiber.Group("/api/v1")
	// user
	userRepo := repository.NewUserRepository(app.DB)
	userService := service.NewUserService(userRepo)
	userHandler := v1.NewUserHandler(userService)
	userHandler.Setup(apiV1)

}
