package app

import (
	"go.uber.org/zap"
	"gopkg.in/telebot.v4"
)

func (a *App) ping(c telebot.Context) error {
	a.logger.Info("/ping",
		zap.Any("storage", a.storage),
	)

	return c.Send(a.makePing())
}
