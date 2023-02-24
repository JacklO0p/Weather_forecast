package listener

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/api/weather"
	"github.com/JacklO0p/weather_forecast/controllers"
	"github.com/JacklO0p/weather_forecast/globals"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	models2 "github.com/JacklO0p/weather_forecast/models"
)

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	command := strings.ToLower(update.Message.Text)

	if !globals.IsProgramStarted {
		if command == "/start" {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Program started, use /help for the list of available commands",
			})
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Must first start the bot using /start",
			})

			return
		}

		globals.IsProgramStarted = true
	}

	if len(command) >= 1 {

		if !strings.Contains(command, " ") {

			if isCommandPresent(command) {
				if command == "/stop" {
					b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: update.Message.Chat.ID,
						Text:   "The program will now end, c'ya",
					})

					ctx.Done()
				}

				if command == "/timeframe" {
					b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: update.Message.Chat.ID,
						Text:   "Current time frame: " + strconv.Itoa(weather.TimeFrame),
					})
				}

				if command == "/help" {
					b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: update.Message.Chat.ID,
						Text:   "***Command list:***\n\n" + ListOfCommands(),
					})
				}

				if command == "/meteo" {
					controllers.GetWeather()
				}

			} else {
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   command + " is not present in the command list\nType /help for a list of all possible commands",
				})
			}

		} else {
			Check := strings.Split(command, " ")

			if isCommandPresent(Check[0]) {
				//the command is present in the list, need to check which one it is now

				//new location command
				if Check[0] == "/location" {
					globals.CurrentLocation = Check[1]

					b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: update.Message.Chat.ID,
						Text:   "Location updated succesfully!\nNow recording " + Check[1],
					})

					user := models2.User{
						ChatID:    update.Message.Chat.ID,
						Location:  Check[1],
						Timeframe: weather.TimeFrame,
					}

					err := models2.UpdateUser(&user)
					if err != nil {
						b.SendMessage(ctx, &bot.SendMessageParams{
							ChatID: update.Message.Chat.ID,
							Text:   "Failed to update location",
						})

					} else {
						b.SendMessage(ctx, &bot.SendMessageParams{
							ChatID: update.Message.Chat.ID,
							Text:   "Location updated succesfully!\nNow recording " + Check[1],
						})
					}
					
				}

				//timeframe command
				if Check[0] == "/newtimeframe" {

					newTimeFrame, err := strconv.Atoi(Check[1])
					if err != nil {
						b.SendMessage(ctx, &bot.SendMessageParams{
							ChatID: update.Message.Chat.ID,
							Text:   "Time frame not valid",
						})

						weather.TimeFrame = 1
					} else {
						b.SendMessage(ctx, &bot.SendMessageParams{
							ChatID: update.Message.Chat.ID,
							Text:   "The previous time frame: " + strconv.Itoa(weather.TimeFrame) + "\nis replaced by the new one: " + Check[1],
						})

						weather.TimeFrame = newTimeFrame
					}
				}

			} else {
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   Check[0] + " is not present in the command list\nType /help for a list of all possible commands",
				})
			}
		}
	}

}

func TelegramListener() {
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
