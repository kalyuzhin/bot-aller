package app

import (
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"
	"gopkg.in/telebot.v4"
)

func (a *App) pingAt(c telebot.Context) error {
	a.logger.Info("/ping_at",
		zap.Any("storage", a.storage),
	)
	msg := c.Text()
	splitted := strings.SplitN(msg, " ", 2)
	if len(splitted) < 2 {
		a.logger.Error("/ping_at")

		return errors.New("invalid arguments")
	}
	date, err := time.Parse("2006-01-02 15:04", splitted[1])
	if err != nil {
		a.logger.Error("/ping", zap.Any("error", err))

		return err
	}
	moscowLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		a.logger.Error("/ping", zap.Any("error", err))

		return err
	}
	dateMoscow := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0,
		moscowLocation)
	go a.workerAt(c, dateMoscow)

	return nil
}
