package app

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"gopkg.in/telebot.v4"
)

func (a *App) pingIn(c telebot.Context) error {
	a.logger.Info("/ping_in",
		zap.Any("storage", a.storage),
	)
	msg := c.Text()
	splitted := strings.Split(msg, " ")
	if len(splitted) < 2 {
		a.logger.Error("/ping_in")

		return errors.New("invalid arguments")
	}
	val, err := strconv.Atoi(splitted[1])
	if err != nil {
		a.logger.Error("/ping", zap.Any("error", err))

		return err
	}
	t := time.Duration(val) * time.Second
	go a.workerIn(c, t)

	return nil
}
