package main

import "github.com/msakp/golang-web-template/internal/app"

//	@title		Golang clean-arch Web Template
//	@version	1.0
//	@host		localhost:3000
//	@BasePath	/api/v1

//	@securitydefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@descrtiption				"access token 'Bearer {token}'"
func main() {
	app := app.NewApp()
	app.Start()
}
