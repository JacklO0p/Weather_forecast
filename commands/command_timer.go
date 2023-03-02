package commands

import (
	"context"
	"fmt"
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
	return "Display the current timer or set a new timer"
}

func (c *CommandTimer) Execute(ctx context.Context, b *bot.Bot, update *models.Update, args []string) error {
	user := models2.User{
		ChatID: update.Message.Chat.ID,
	}

	globals.Db.Where(&user).First(&user)

	currTimer := user.Timer

	if len(args) != 0 {
		newTimer, err := strconv.Atoi(args[0])
		
		if err != nil {
			fmt.Print("Error: ", err)
		}

		user.Timer = newTimer
		globals.Db.Save(&user)

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "New timer: " + strconv.Itoa(newTimer),
		})

		globals.Timer = newTimer

	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Current timer: " + strconv.Itoa(currTimer),
		})

		globals.Timer = currTimer
	}
	

	return nil
}
