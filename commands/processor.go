package commands

import (
	"context"
	"errors"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandProcessor struct {
	commands []Commads
}

func (c *CommandProcessor) AddCommand(command Commads) {
	c.commands = append(c.commands, command)
}

func (c *CommandProcessor) Process(ctx context.Context, b *bot.Bot, update *models.Update) error {
	command := strings.ToLower(update.Message.Text)

	if len(command) >= 1 {
		check := strings.Split(command, " ")

		for _, cmd := range c.commands {
			if check[0] == cmd.Command() {
				return cmd.Execute(b, update, check[1:])
			}
		}

	}

	return errors.New("command not found")
}
