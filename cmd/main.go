package main

import "github.com/msakp/golang-web-template/internal/app"

//	@title		Golang clean-arch Web Template
//	@version	1.0
//  @host localhost:8080
//  @BasePath /api/v1

func main() {
	app := app.NewApp()
	app.Start()
}
