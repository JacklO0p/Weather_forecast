package commands

import (
	"context"
	"strconv"
	"time"

	"github.com/JacklO0p/weather_forecast/api/location"
	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/globals"
	models2 "github.com/JacklO0p/weather_forecast/models"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandReport struct{}

func (c *CommandReport) Command() string {
	return "/report"
}

func (c *CommandReport) Description() string {
	return "Receive a meteo report of your selected location"
}

func (c *CommandReport) Execute(ctx context.Context, b *bot.Bot, update *models.Update, args []string) error {
	meteoReport(ctx, b, update)

	return nil
}

func meteoReport(ctx context.Context, b *bot.Bot, update *models.Update) {
	user := models2.User{}

	user.ChatID = update.Message.Chat.ID
	globals.Db.Where(&user).First(&user)

	if user.Location != "" {

		if location.IsValid(user.Location) {

			msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Fetching weather data...",
			})

			if err == nil {

				hours := time.Duration(globals.Timer) * time.Minute

				go func() {
					b.EditMessageText(ctx, &bot.EditMessageTextParams{
						ChatID:    update.Message.Chat.ID,
						MessageID: msg.ID,
						Text:      telegram.GetReport(user.Location) + user.Location + "\n\nThe next report will be sent in " + strconv.Itoa(int(hours.Hours())) + " hours and " + strconv.Itoa(int(hours.Minutes())%60) + " minutes",
					})
				}()
			}
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "The location you set is not valid, type /newlocation <location> to set a new one",
			})
		}

	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "No location set, type /location <location> to set one",
		})
	}
}
