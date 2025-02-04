package main

import "github.com/msakp/golang-web-template/internal/app"

func main() {
	app := app.NewApp()
	app.Start()
}
