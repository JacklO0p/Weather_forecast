package listener

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/globals"
	models2 "github.com/JacklO0p/weather_forecast/models"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

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

				if command == "/timer" {
					SendMessage(ctx, b, update, "Current timer: "+strconv.Itoa(globals.TimeFrame)+" minutes")
				}

				if command == "/help" {
					SendMessage(ctx, b, update, "List of all possible commands:\n"+ListOfCommands())
				}

				if command == "/report" {
					SendMeteoReport(ctx, b, update)

				}

			} else {
				SendMessage(ctx, b, update, "Command not found, type /help for a list of all possible commands")
			}

		} else {
			Check := strings.Split(command, " ")

			if isCommandPresent(Check[0]) {

				if Check[0] == "/location" {
					location := strings.Join(Check[1:], " ")

					SendMessage(ctx, b, update, "Location updated succesfully!\nNow recording "+location)

					user := models2.User{
						ChatID:   update.Message.Chat.ID,
						Location: location,
						Timer:    globals.TimeFrame,
					}

					err := models2.UpdateUser(&user)
					if err != nil {
						SendMessage(ctx, b, update, "Error while updating the user")

					}

				}

				if Check[0] == "/newtimer" {

					newTimer, err := strconv.Atoi(Check[1])

					if err != nil {
						SendMessage(ctx, b, update, "Couldn't update timer\nIt will now be set to 240 minute")
					} else {
						SendMessage(ctx, b, update, "Timer updated succesfully!\nIt will now be set to "+Check[1]+" minutes")

						user := models2.User{
							ChatID: update.Message.Chat.ID,
							Timer:  newTimer,
						}

						err = models2.UpdateUser(&user)
						if err != nil {
							SendMessage(ctx, b, update, "Error while updating the user")

						}
					}

				}

			} else {
				SendMessage(ctx, b, update, "Command not found, type /help for a list of all possible commands")
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

	globals.Bot = b

	b.Start(ctx)
}

func SendMessage(ctx context.Context, b *bot.Bot, update *models.Update, message string) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   message,
	})
}

func SendMeteoReport(ctx context.Context, b *bot.Bot, update *models.Update) {
	user := models2.User{}

	user.ChatID = update.Message.Chat.ID
	globals.Db.Where(&user).First(&user)

	if user.Location != "" {
		msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Fetching weather data...",
		})
		if err == nil {
			go func() {
				b.EditMessageText(ctx, &bot.EditMessageTextParams{
					ChatID:    update.Message.Chat.ID,
					MessageID: msg.ID,
					Text:      telegram.GetReport(user.Location) + user.Location,
				})
			}()
		}
	} else {
		SendMessage(ctx, b, update, "No location set, type /location <location> to set one")
	}
}
