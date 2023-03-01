package commands

import (
	"context"
	"fmt"
	"strings"

	models2 "github.com/JacklO0p/weather_forecast/models"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandNewLocation struct{}

func (c *CommandNewLocation) Command() string {
	return "/newlocation"
}

func (c *CommandNewLocation) Description() string {
	return "Switch location"
}

func (c *CommandNewLocation) Execute(ctx context.Context, b *bot.Bot, update *models.Update, args []string) error {

	user := models2.User{
		ChatID:   update.Message.Chat.ID,
		Location: strings.Join(args, " "),
	}

	err := models2.UpdateUser(&user)
	if err != nil {
		fmt.Println("Error while updating location: ", err)

		return err
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "New Location: " + user.Location,
	})

	return nil
}
