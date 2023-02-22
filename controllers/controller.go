package controllers

import (
	"fmt"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/api/weather"
)

func GetWeather() {
	var result map[string]interface{}

	result = weather.GetWeatherFromLocation()
	if result == nil {
		fmt.Print("Error while getting the value, result is nil\n")

		telegram.SendTelegramMessage("Found error in the code,\n\nMake sure the location given is an existing city, the location will now be set to Trevignano", telegram.CHATID)
		telegram.CurrentLocation = "Trevignano"
	}

	dividedMessage := telegram.DivideMessages(result)

	if dividedMessage == "" {
		telegram.SendTelegramMessage("no raining today or tomorrow: "+weather.CurrentDateString+"    /    "+weather.TomorrowDateString+"\nLocation: "+telegram.CurrentLocation, telegram.CHATID)
	} else {
		telegram.SendTelegramMessage(dividedMessage+"\n\nLocation: "+telegram.CurrentLocation, telegram.CHATID)
	}

}
