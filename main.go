package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/api/telegram/listener"
	"github.com/JacklO0p/weather_forecast/api/weather"
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

	duration := time.Duration(weather.TimeFrame) * time.Minute

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			log.Println("mannaggia cristo")
			if duration != time.Duration(weather.TimeFrame)*time.Minute {
				duration = time.Duration(weather.TimeFrame) * time.Minute
				ticker.Stop()
			}

			ticker = time.NewTicker(duration)

			var user []models.User
			globals.Db.Find(&user)

			for _, u := range user {

				if u.Location != "" && u.SendMessage {

					listener.Bot.SendMessage(context.Background(), &bot.SendMessageParams{
						ChatID: u.ChatID,
						Text:   telegram.GetReport(u.Location) + u.Location,
					})

				}

			}

			fmt.Print("\nReport sent successfully\n")
		}
	}


		

}


