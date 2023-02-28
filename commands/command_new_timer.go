package commands

import (
	"context"
	"fmt"
	"strconv"

	models2 "github.com/JacklO0p/weather_forecast/models"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandNewTimer struct{}

func (c *CommandNewTimer) Command() string {
	return "/newtimer"
}

func (c *CommandNewTimer) Description() string {
	return "Set a new timer for your reports"
}

func (c *CommandNewTimer) Execute(ctx context.Context, b *bot.Bot, update *models.Update, args []string) error {

	argsToInt, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error while converting args to int: ", err)
		return err
	}

	user := models2.User{
		ChatID: update.Message.Chat.ID,
		Timer:  argsToInt,
	}

	models2.UpdateUser(&user)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "New timer: " + strconv.Itoa(user.Timer),
	})

	return nil
}
