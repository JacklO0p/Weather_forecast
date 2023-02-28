package commands

import (
	"context"

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

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "Program started, /help for the list of all possible commands",
	})

	return nil
}
