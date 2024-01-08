package main

import (
	appconfig "github.com/alichz2001/fidibo/internal/config"
	"github.com/alichz2001/fidibo/internal/repository"
	"github.com/alichz2001/fidibo/internal/router"
	"github.com/alichz2001/fidibo/internal/server"
	"github.com/alichz2001/fidibo/internal/service"
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	appCfg := appconfig.LoadConfig(appconfig.LoadConfigOptions{RunRedis: true, RunMysql: true, PanicByErr: true})

	appCfg.Logger = logger

	appRepo := repository.InitRepo(appCfg)
	bookRepo := repository.NewBookRepository(appCfg, appRepo)
	bookService := service.InitBookService(appCfg, bookRepo)

	app := server.NewWebServer(appCfg)

	router.AddBookRoutes(app, bookService)

	app.Run(appCfg.GetServerAddress())
}
