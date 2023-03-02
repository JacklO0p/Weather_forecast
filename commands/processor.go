package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/JacklO0p/weather_forecast/globals"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandProcessor struct {
	commands []Commads
}

func (c *CommandProcessor) AddCommand(command Commads) {
	fmt.Println("Added command: " + command.Command())

	c.commands = append(c.commands, command)
}

func (c *CommandProcessor) Process(ctx context.Context, b *bot.Bot, update *models.Update) error {
	command := strings.ToLower(update.Message.Text)

	if len(command) >= 1 {
		check := strings.Split(command, " ")
	fmt.Println("Comando non presente: " + check[0])

		if check[0] == "/help" {
			var list string
			
			for _, cmd := range c.commands {
				list += "- " + cmd.Command() + "   |   " + cmd.Description() + "\n"
			}

				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   "Command list:\n" + list,
				})

				return nil
			}

			for _, cmd := range c.commands {
				if check[0] == cmd.Command() {
					return cmd.Execute(ctx, b, update, check[1:])
				}
			}

	}


	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Comando non presentte, /help per lista dei comandi",
	})

	return nil
}
