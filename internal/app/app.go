package app

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v4"

	"github.com/kalyuzhin/bot-aller/pkg/config"
)

var storage map[string]struct{}

// Run ...
func Run(conf config.Config) error {
	settings := telebot.Settings{
		Token: conf.Token,
	}
	storage = make(map[string]struct{})
	bot, err := telebot.NewBot(settings)
	if err != nil {
		return err
	}
	bot.Use(MiddleWare)
	bot.Handle("/ping", func(c telebot.Context) error {
		sb := strings.Builder{}
		for username, _ := range storage {
			sb.WriteString(fmt.Sprintf("@%s ", username))
		}

		return c.Send(sb.String())
	})
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		return nil
	})
	bot.Start()

	return nil
}

func MiddleWare(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		updateStorage(c.Sender().Username)

		return next(c)
	}
}

func updateStorage(username string) {
	if _, ok := storage[username]; !ok {
		storage[username] = struct{}{}
	}
}
