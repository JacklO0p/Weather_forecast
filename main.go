package main

import (
	"fmt"
	"time"

	"github.com/JacklO0p/weather_forecast/api/telegram/listener"
	"github.com/JacklO0p/weather_forecast/api/weather"
	"github.com/JacklO0p/weather_forecast/controllers"
)

var Check []string

func main() {
	var beginning int = 1

	//start telegram bot listener
	listener.Inizializer()
	go listener.TelegramListener()

    

	if listener.CheckToStart || beginning == 1 {
		beginning = 2

		duration := time.Duration(weather.TimeFrame) * time.Minute

		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		for {
			select {
				
			case <-ticker.C:

				if duration != time.Duration(weather.TimeFrame) * time.Minute {
					duration = time.Duration(weather.TimeFrame) * time.Minute
					ticker.Stop()
				}

				ticker = time.NewTicker(duration)

				controllers.GetWeather()
				fmt.Print("\nReport sent successfully\n")
			}
		}
	}
	

}
