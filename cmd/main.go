package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/msakp/golang-web-template/domain/dto"
	"github.com/msakp/golang-web-template/internal/config"
	"github.com/msakp/golang-web-template/internal/infrastructure/database"
	"github.com/msakp/golang-web-template/internal/repository"
)

func main() {
	config := config.New()
	db, err := database.NewPg(config)
	if err != nil {
		log.Fatalf("PIZDEC DB: %s", err.Error())
	}
	defer db.CloseConn()
	err = db.Migrate()
	if err != nil{
		log.Fatalf("PIZDEC MIGRACIYAM: %s", err.Error())
	}
	repo := repository.NewUserRepository(db)
	u1 := dto.UserRegister{
		Name: "msa",
		Email: "arst@arst",
		PasswordUnhashed: "qwerty",
	}
	err = repo.Create(&u1)
	if err != nil{
		log.Fatalf("PIZDEC REPE: %s", err.Error())
	}
	app := fiber.New()
	v1 := app.Group("/api/v1")
	v1.Get("/ping", defaultHandler)
	log.Fatal(app.Listen("0.0.0.0:8080"))

}

func defaultHandler(ctx fiber.Ctx) error {
	return ctx.Status(200).JSON("pong")
}
