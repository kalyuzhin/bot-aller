package app

import (
	"gopkg.in/telebot.v4"

	"github.com/kalyuzhin/bot-aller/pkg/config"
)

// Run ...
func Run(conf config.Config) error {
	settings := telebot.Settings{
		Token: conf.Token,
	}
	bot, err := telebot.NewBot(settings)
	if err != nil {
		return err
	}
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send(c.Chat().Username)
	})
	bot.Start()

	return nil
}
