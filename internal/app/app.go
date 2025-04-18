package app

import (
	"fmt"
	"log"
	"strings"

	"go.uber.org/zap"
	"gopkg.in/telebot.v4"

	"github.com/kalyuzhin/bot-aller/pkg/config"
)

var storage map[string]struct{}

var logger *zap.Logger

var bot *telebot.Bot

// Run ...
func Run(conf config.Config) error {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalf("unable to initialize zap logger: %v", err)
	}
	defer logger.Sync()

	settings := telebot.Settings{
		Token: conf.Token,
	}
	storage = make(map[string]struct{})
	bot, err = telebot.NewBot(settings)
	if err != nil {
		return err
	}
	bot.Use(MiddleWare)
	bot.Handle("/ping", func(c telebot.Context) error {
		logger.Info("/ping",
			zap.Any("storage", storage),
		)
		sb := strings.Builder{}
		for username, _ := range storage {
			sb.WriteString(fmt.Sprintf("@%s ", username))
		}

		return c.Send(sb.String())
	})
	bot.Handle(telebot.OnUserJoined, func(c telebot.Context) error {
		return nil
	})
	bot.Handle(telebot.OnVoice, func(c telebot.Context) error {
		return nil
	})
	bot.Handle(telebot.OnVideo, func(c telebot.Context) error {
		return nil
	})
	bot.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		return nil
	})
	bot.Handle(telebot.OnSticker, func(c telebot.Context) error {
		return nil
	})
	bot.Handle(telebot.OnAudio, func(c telebot.Context) error {
		return nil
	})
	bot.Start()

	return nil
}

func MiddleWare(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		name := c.Sender().Username
		if _, ok := storage[name]; ok {
			return next(c)
		}
		if name == "" {
			logger.Info("User will not be mentioned",
				zap.String("first name", c.Sender().FirstName),
				zap.String("first name", c.Sender().LastName),
			)
			return next(c)
		}
		storage[name] = struct{}{}
		logger.Info("On user joined",
			zap.String("username", name),
			zap.Any("storage", storage),
		)

		return next(c)
	}
}
