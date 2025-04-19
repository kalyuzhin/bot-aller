package app

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gopkg.in/telebot.v4"

	"github.com/kalyuzhin/bot-aller/pkg/config"
)

// App
type App struct {
	bot     *telebot.Bot
	logger  *zap.Logger
	conf    *config.Config
	storage map[string]struct{}
}

// NewApp creates a new instance of App
func NewApp(logger *zap.Logger, conf *config.Config) (*App, error) {
	settings := telebot.Settings{
		Token: conf.Token,
	}
	storage := make(map[string]struct{}, 10000)
	bot, err := telebot.NewBot(settings)
	if err != nil {
		return nil, err
	}

	return &App{
		bot:     bot,
		logger:  logger,
		conf:    conf,
		storage: storage,
	}, nil
}

// Run is an action to setup middlewares, handlers and start up bot
func (a *App) Run() {
	a.bot.Use(a.MiddleWare)
	a.SetupHandlers()

	a.bot.Start()
}

// Middleware ...
func (a *App) MiddleWare(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		name := c.Sender().Username
		if _, ok := a.storage[name]; ok {
			return next(c)
		}
		if name == "" {
			a.logger.Info("User will not be mentioned",
				zap.String("first name", c.Sender().FirstName),
				zap.String("first name", c.Sender().LastName),
			)

			return next(c)
		}
		a.storage[name] = struct{}{}
		a.logger.Info("On user joined",
			zap.String("username", name),
			zap.Any("storage", a.storage),
		)

		return next(c)
	}
}

func (a *App) workerIn(ctx telebot.Context, t time.Duration) {
	ticker := time.NewTicker(t)
	<-ticker.C
	err := ctx.Send(a.makePing())
	if err != nil {
		a.logger.Error("unable to send ping", zap.Error(err))
	}
}

func (a *App) workerAt(ctx telebot.Context, t time.Time) {
	pingAt := time.Until(t)
	ticker := time.NewTicker(pingAt)
	<-ticker.C
	err := ctx.Send(a.makePing())
	if err != nil {
		a.logger.Error("unable to send ping", zap.Error(err))
	}
}

func (a *App) makePing() string {
	sb := strings.Builder{}
	for username := range a.storage {
		sb.WriteString(fmt.Sprintf("@%s ", username))
	}

	return sb.String()
}
