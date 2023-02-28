package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Commads interface {
	Command() string
	Execute(ctx context.Context, b *bot.Bot, update *models.Update, args []string) error
	Description() string
}
