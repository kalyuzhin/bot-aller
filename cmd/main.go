package main

import (
	"github.com/kalyuzhin/bot-aller/internal/app"
	"log"

	"github.com/kalyuzhin/bot-aller/pkg/config"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = app.Run(*conf)
	if err != nil {
		log.Fatal(err)
	}
}
