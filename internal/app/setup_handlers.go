package app

import (
	"gopkg.in/telebot.v4"
)

// SetupHandlers ...
func (a *App) SetupHandlers() {
	a.bot.Handle("/ping", a.ping)
	a.bot.Handle("/ping_in", a.pingIn)
	a.bot.Handle("/ping_at", a.pingAt)
	a.bot.Handle(telebot.OnText, nop)
	a.bot.Handle(telebot.OnUserJoined, nop)
	a.bot.Handle(telebot.OnVoice, nop)
	a.bot.Handle(telebot.OnVideo, nop)
	a.bot.Handle(telebot.OnPhoto, nop)
	a.bot.Handle(telebot.OnSticker, nop)
	a.bot.Handle(telebot.OnAudio, nop)
}
