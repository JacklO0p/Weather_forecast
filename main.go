package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/api/telegram/listener"
	"github.com/JacklO0p/weather_forecast/globals"
	"github.com/JacklO0p/weather_forecast/models"
	"github.com/JacklO0p/weather_forecast/utils"
	"github.com/go-telegram/bot"
)

func main() {

	//connect to database
	utils.Connect()
	utils.MigrateDB()

	listener.Inizializer()
	go listener.TelegramListener()

	duration := time.Duration(globals.Timer) * time.Minute

	fmt.Print("duration: ", duration, "\n")

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			if duration != time.Duration(globals.TimerDuration) * time.Minute {
				duration = time.Duration(globals.TimerDuration) * time.Minute
				ticker.Stop()
			}


			ticker = time.NewTicker(duration)

			var user []models.User
			globals.Db.Find(&user)

			for _, u := range user {

				if u.Location != "" && u.SendMessage  && globals.IsProgramStarted {
					
					globals.TimerDuration = time.Duration(u.Timer) * time.Minute

					globals.Bot.SendMessage(context.Background(), &bot.SendMessageParams{
						ChatID: u.ChatID,
						Text:   telegram.GetReport(u.Location) + "\n\nLocation: " + u.Location + "\n\nNext report in " + globals.TimerDuration.String(),
					})

				}

			}

			fmt.Print("\nReport sent successfully\n")
		}
	}

}
