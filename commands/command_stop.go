package commands

import (
	"context"
	"fmt"

	"github.com/JacklO0p/weather_forecast/globals"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandStop struct{}

func (c *CommandStop) Command() string {
	return "/stop"
}

func (c *CommandStop) Description() string {
	return "Stop the program"
}

func (c *CommandStop) Execute(ctx context.Context, b *bot.Bot, update *models.Update, args []string) error {

	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "stopping the program...",
	})
	
	if err == nil {
		go func() {
			b.EditMessageText(ctx, &bot.EditMessageTextParams{
				ChatID:    update.Message.Chat.ID,
				MessageID: msg.ID,
				Text:      "program stopped",
			})
		}()
	} else {
		fmt.Print(err)
		return err
	}

	globals.IsProgramStarted = false

	return nil
}
