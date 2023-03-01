package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/JacklO0p/weather_forecast/api/location"
	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/api/telegram/listener"
	"github.com/JacklO0p/weather_forecast/globals"
	models2 "github.com/JacklO0p/weather_forecast/models"
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

	fmt.Print("duration: ", duration)
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:

			if duration != time.Duration(globals.Timer)*time.Minute {
				duration = time.Duration(globals.Timer) * time.Minute
				ticker.Stop()
			}

			fmt.Print("new duration: ", duration)

			ticker = time.NewTicker(duration)

			var user []models2.User
			globals.Db.Find(&user)

			for _, u := range user {

				if u.Location != "" && u.SendMessage {

					if location.IsValid(u.Location) {
						globals.Bot.SendMessage(context.Background(), &bot.SendMessageParams{
							ChatID: u.ChatID,
							Text:   telegram.GetReport(u.Location) + u.Location + "\n\nThe next report will be sent in " + strconv.Itoa(int(duration.Hours())) + " hours and " + strconv.Itoa(int(duration.Minutes())%60) + " minutes",
						})
					} else {
						globals.Bot.SendMessage(context.Background(), &bot.SendMessageParams{
							ChatID: u.ChatID,
							Text:   "The location you set is not valid, type /newlocation <location> to set a new one\n\nThe next report will be sent in " + strconv.Itoa(int(duration.Hours())) + " hours and " + strconv.Itoa(int(duration.Minutes())%60) + " minutes",
						})
					}

				}

			}

			fmt.Print("\nReport sent successfully\n")
		}
	}

}
