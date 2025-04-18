package app

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

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

		return c.Send(makePing())
	})
	bot.Handle("/ping_in", func(c telebot.Context) error {
		logger.Info("/ping_in",
			zap.Any("storage", storage),
		)
		msg := c.Text()
		splitted := strings.Split(msg, " ")
		if len(splitted) < 2 {
			logger.Error("/ping_in")

			return errors.New("invalid arguments")
		}
		val, err := strconv.Atoi(splitted[1])
		if err != nil {
			logger.Error("/ping", zap.Any("error", err))

			return err
		}
		t := time.Duration(val) * time.Minute
		go workerIn(c, t)

		return nil

	})
	bot.Handle("/ping_at", func(c telebot.Context) error {
		logger.Info("/ping_at",
			zap.Any("storage", storage),
		)
		msg := c.Text()
		splitted := strings.SplitN(msg, " ", 2)
		if len(splitted) < 2 {
			logger.Error("/ping_at")

			return errors.New("invalid arguments")
		}
		date, err := time.Parse("2006-01-02 15:04", splitted[1])
		if err != nil {
			logger.Error("/ping", zap.Any("error", err))

			return err
		}
		moscowLocation, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			logger.Error("/ping", zap.Any("error", err))

			return err
		}
		dateMoscow := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0,
			moscowLocation)
		go workerAt(c, dateMoscow)

		return nil
	})
	bot.Handle(telebot.OnText, func(c telebot.Context) error { return nil })
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

func workerIn(ctx telebot.Context, t time.Duration) {
	ticker := time.NewTicker(t)
	for {
		select {
		case <-ticker.C:
			err := ctx.Send(makePing())
			if err != nil {
				logger.Error("unable to send ping", zap.Error(err))
			}

			return
		}
	}
}

func workerAt(ctx telebot.Context, t time.Time) {
	duration := time.Until(t)
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
			err := ctx.Send(makePing())
			if err != nil {
				logger.Error("unable to send ping", zap.Error(err))
			}

			return
		}
	}
}

func makePing() string {
	sb := strings.Builder{}
	for username, _ := range storage {
		sb.WriteString(fmt.Sprintf("@%s ", username))
	}

	return sb.String()
}
