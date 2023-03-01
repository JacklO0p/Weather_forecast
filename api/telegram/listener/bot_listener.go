package listener

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/commands"
	"github.com/JacklO0p/weather_forecast/globals"
	models2 "github.com/JacklO0p/weather_forecast/models"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var cmds *commands.CommandProcessor = &commands.CommandProcessor{}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	err := cmds.Process(ctx, b, update)
	if err != nil {
		fmt.Println(err)
	}
}

func TelegramListener() {

	cmds.AddCommand(&commands.CommandStart{})
	cmds.AddCommand(&commands.CommandTimer{})
	cmds.AddCommand(&commands.CommandNewTimer{})
	cmds.AddCommand(&commands.CommandReport{})
	cmds.AddCommand(&commands.CommandNewLocation{})
	cmds.AddCommand(&commands.CommandLocation{})
	cmds.AddCommand(&commands.CommandStop{})
	
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
