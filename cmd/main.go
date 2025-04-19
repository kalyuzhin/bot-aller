package main

import (
	"log"

	"go.uber.org/zap"

	app_pkg "github.com/kalyuzhin/bot-aller/internal/app"
	"github.com/kalyuzhin/bot-aller/pkg/config"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("unable to initialize zap logger: %v", err)
	}
	defer logger.Sync()

	app, err := app_pkg.NewApp(logger, conf)
	if err != nil {
		logger.Fatal("unable to initialize app", zap.Error(err))
	}

	app.Run()
}
