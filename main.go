package main

import (
	"fmt"
	"time"

	"github.com/JacklO0p/weather_forecast/api/telegram/listener"
	"github.com/JacklO0p/weather_forecast/api/weather"
	"github.com/JacklO0p/weather_forecast/utils"
)

func main() {

	//connect to database
	utils.Connect()
	utils.MigrateDB()

	//start telegram bot listener
	listener.Inizializer()
	go listener.TelegramListener(0)

	duration := time.Duration(weather.TimeFrame) * time.Minute

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:

			if duration != time.Duration(weather.TimeFrame)*time.Minute {
				duration = time.Duration(weather.TimeFrame) * time.Minute
				ticker.Stop()
			}

			ticker = time.NewTicker(duration)

			listener.TelegramListener(1)
			fmt.Print("\nReport sent successfully\n")
		}
	}

}
