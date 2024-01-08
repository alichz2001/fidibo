package server

import (
	"fmt"
	"github.com/alichz2001/fidibo/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
	"os/signal"
)

type WebApp struct {
	*fiber.App
}

func NewWebServer(appCfg *config.AppConfig) (app *WebApp) {

	f := fiber.New()
	f.Use(requestid.New())
	f.Use(logger.New())

	app = &WebApp{f}

	return
}

func (app *WebApp) Run(addr string) {

	// graceful shutdown by interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()
	//start listener
	if err := app.Listen(addr); err != nil {
		log.Panic(err)
	}
}
