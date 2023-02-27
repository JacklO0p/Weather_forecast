package commands

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Commads interface {
	Command() string
	Execute(bot *bot.Bot, update *models.Update, args []string) error
}

func (c *CommandNewLocation) Execute(b *bot.Bot, update *models.Update, args []string) error {
	return nil
}
