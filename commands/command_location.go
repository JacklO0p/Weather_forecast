package commands

import (
	"context"

	"github.com/JacklO0p/weather_forecast/globals"
	models2 "github.com/JacklO0p/weather_forecast/models"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandLocation struct{}

func (c *CommandLocation) Command() string {
	return "/location"
}

func (c *CommandLocation) Description() string {
	return "Display the current location"
}

func (c *CommandLocation) Execute(ctx context.Context, b *bot.Bot, update *models.Update, args []string) error {
	user := models2.User{
		ChatID: update.Message.Chat.ID,
	}

	globals.Db.Where(&user).First(&user)

	loc := user.Location

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Current location: " + loc,
	})

	return nil
}
