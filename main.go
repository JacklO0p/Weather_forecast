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

	//start telegram bot listener
	globals.IsProgramStarted = false
	listener.Inizializer()
	go listener.TelegramListener()

	duration := time.Duration(globals.Timer) * time.Minute

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			if duration != time.Duration(globals.Timer) * time.Minute {
				duration = time.Duration(globals.Timer) * time.Minute
				ticker.Stop()
			}

			fmt.Print("new duration: ", duration)


			ticker = time.NewTicker(duration)

			if globals.IsProgramStarted {
				var user []models.User
				globals.Db.Find(&user)
	
				for _, u := range user {
	
					if u.Location != "" && u.SendMessage {

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

}
