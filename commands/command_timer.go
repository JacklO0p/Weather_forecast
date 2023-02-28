package commands

import (
	"context"
	"strconv"

	"github.com/JacklO0p/weather_forecast/globals"
	models2 "github.com/JacklO0p/weather_forecast/models"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandTimer struct{}

func (c *CommandTimer) Command() string {
	return "/timer"
}

func (c *CommandTimer) Description() string {
	return "Display the current timer"
}

func (c *CommandTimer) Execute(ctx context.Context, b *bot.Bot, update *models.Update, args []string) error {
	user := models2.User{
		ChatID: update.Message.Chat.ID,
	}

	globals.Db.Where(&user).First(&user)

	currTimer := user.Timer

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Current timer: " + strconv.Itoa(currTimer),
	})

	return nil
}
