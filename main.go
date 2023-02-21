package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/routes"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/labstack/echo/v4"
)

var Check []string

func main() {

	//start telegram bot listener
	go telegramListener()

	e := echo.New()

	routes.RegisteredRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	command := update.Message.Text

	if len(command) >= 2 {
		Check := strings.Split(command, " ")

		if Check[0] == "/location" {
			telegram.CurrentLocation = Check[1]

			fmt.Print(telegram.CurrentLocation)

			b.SendMessage(ctx, &bot.SendMessageParams{

				ChatID: update.Message.Chat.ID,
				Text:   "Location updated succesfully!\nNow recording " + Check[1],
			})
		}
	}

}

func telegramListener() {
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
