package commands

import (
	"context"

	"github.com/JacklO0p/weather_forecast/globals"
	models2 "github.com/JacklO0p/weather_forecast/models"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandStart struct{}

func (c *CommandStart) Command() string {
	return "/start"
}

func (c *CommandStart) Description() string {
	return "Start the program"
}

func (c *CommandStart) Execute(ctx context.Context, b *bot.Bot, update *models.Update, args []string) error {

	user := models2.User{
		ChatID: update.Message.Chat.ID,
	}

	globals.Db.Where(&user).First(&user)
	
	globals.Timer = user.Timer

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Program started, /help for the list of all possible commands",
	})

	globals.IsProgramStarted = true
	
	return nil
}
