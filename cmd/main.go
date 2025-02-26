package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/msakp/golang-web-template/internal/app"
	"github.com/msakp/golang-web-template/pkg/logger"
)

//	@title		Golang clean-arch Web Template
//	@version	1.0
//	@host		localhost:808
//	@BasePath	/api/v1

// @securitydefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @descrtiption				"access token 'Bearer {token}'"
func main() {
	var (
		shutDownGroup sync.WaitGroup
		ctx           = context.Background()
		signalCh      = make(chan os.Signal, 1)
	)
	signal.Notify(signalCh, os.Interrupt)

	ctx = logger.CtxWithLogger(ctx)
	app := app.NewApp(ctx)
	// graceful shutdown handler
	go InterruptHandler(app, signalCh, &shutDownGroup)
	app.Start()
	// connections closing on shutdown
	app.CloseConnections(ctx)
}

func InterruptHandler(app *app.App, signalCh chan os.Signal, group *sync.WaitGroup) {
	<-signalCh
	fmt.Printf("\n**GRACEFULLY SHUTTING DOWN**\n\n")
	group.Add(1)
	defer group.Done()
	app.Fiber.ShutdownWithTimeout(15 * time.Second)

}
