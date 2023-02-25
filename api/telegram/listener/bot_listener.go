package listener

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/api/weather"
	"github.com/JacklO0p/weather_forecast/globals"
	models2 "github.com/JacklO0p/weather_forecast/models"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var SendMeteo bool = false

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if SendMeteo && globals.IsProgramStarted {
		result := weather.GetWeatherFromLocation()
		if result == nil {
			fmt.Print("Error while getting the value, result is nil, ")

			SendMessage(ctx, b, update, "Error while getting the value, result is nil, ")

			globals.CurrentLocation = "Trevignano"
			return
		}

		dividedMessage := telegram.DivideMessage(result)

		if dividedMessage == "" {
			SendMessage(ctx, b, update, "No raining tomorrow\nLocation: "+globals.CurrentLocation)
		}

		SendMeteo = false
	}

	command := strings.ToLower(update.Message.Text)

	if !globals.IsProgramStarted {
		if command == "/start" {
			SendMessage(ctx, b, update, "The program has started, type /help for a list of all possible commands")
		} else {
			SendMessage(ctx, b, update, "The program has not started yet, type /start to start it")

			return
		}

		globals.IsProgramStarted = true
	}

	if len(command) >= 1 {

		if !strings.Contains(command, " ") {

			if isCommandPresent(command) {
				if command == "/stop" {
					SendMessage(ctx, b, update, "The program has stopped, type /start to start it again")
				}

				if command == "/timeframe" {
					SendMessage(ctx, b, update, "The time frame is set to: "+strconv.Itoa(weather.TimeFrame)+" minutes")
				}

				if command == "/help" {
					SendMessage(ctx, b, update, "List of all possible commands:\n"+ListOfCommands())
				}

				if command == "/meteo" {
					TelegramListener(1)
				}

			} else {
				SendMessage(ctx, b, update, "Command not found, type /help for a list of all possible commands")
			}

		} else {
			Check := strings.Split(command, " ")

			if isCommandPresent(Check[0]) {
				//the command is present in the list, need to check which one it is now

				//new location command
				if Check[0] == "/location" {
					globals.CurrentLocation = Check[1]

					SendMessage(ctx, b, update, "Location updated succesfully!\nNow recording "+Check[1])

					user := models2.User{
						ChatID:    update.Message.Chat.ID,
						Location:  Check[1],
						Timeframe: weather.TimeFrame,
					}

					err := models2.UpdateUser(&user)
					if err != nil {
						SendMessage(ctx, b, update, "Error while updating the user")

					} else {
						SendMessage(ctx, b, update, "User updated succesfully!\nNow recording "+Check[1])
					}

				}

				//timeframe command
				if Check[0] == "/newTimer" {

					newTimer, err := strconv.Atoi(Check[1])
					if err != nil {
						SendMessage(ctx, b, update, "Couldn't update timer\nIt will now be set to 1 minute")

						weather.TimeFrame = 1
					} else {
						SendMessage(ctx, b, update, "Timer updated succesfully!\nIt will now be set to "+Check[1]+" minutes")

						weather.TimeFrame = newTimer
					}
				}

			} else {
				SendMessage(ctx, b, update, "Command not found, type /help for a list of all possible commands")
			}
		}

	}

}

func TelegramListener(meteo int) {

	if meteo == 1 {
		SendMeteo = true
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(telegram.TOKEN, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func SendMessage(ctx context.Context, b *bot.Bot, update *models.Update, message string) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   message,
	})
}
