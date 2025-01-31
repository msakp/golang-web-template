package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/msakp/golang-web-template/internal/infrastructure/config"
	"github.com/msakp/golang-web-template/internal/infrastructure/database"
	"github.com/msakp/golang-web-template/internal/infrastructure/database/sqlc/storage"
)

func main() {
	config := config.New()
	log.Println(config.PostgresUrl)
	db, err := database.NewPg(config)
	if err != nil {
		log.Fatalf("PIZDEC DB: %s", err.Error())
	}
	err = db.CreateUser(db.Ctx, storage.CreateUserParams{Name: "John", Email: "pidor@hui.blt", Password: "sosi moi hui"})
	if err != nil {
		log.Fatalf("PIZDEC SQLC: %s", err.Error())
	}
	app := fiber.New()
	v1 := app.Group("/api/v1")
	v1.Get("/ping", defaultHandler)
	log.Fatal(app.Listen("0.0.0.0:8080"))

}

func defaultHandler(ctx fiber.Ctx) error {
	return ctx.Status(200).JSON("pong")
}
