package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)


func main(){
	app := fiber.New()
	v1 := app.Group("/api/v1")
	v1.Get("/ping", defaultHandler)
	log.Fatal(app.Listen("0.0.0.0:8080"))

}



func defaultHandler(ctx fiber.Ctx)error{
	return ctx.Status(200).JSON("pong")
}



