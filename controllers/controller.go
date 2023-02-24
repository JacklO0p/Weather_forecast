package controllers

import (
	"fmt"
	"strconv"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/api/weather"
	"github.com/JacklO0p/weather_forecast/globals"
	"github.com/JacklO0p/weather_forecast/models"
)

func GetWeather() {
	var result map[string]interface{}

	result = weather.GetWeatherFromLocation()
	if result == nil {
		fmt.Print("Error while getting the value, result is nil\n")

		telegram.SendTelegramMessage("Found error in the code,\n\nMake sure the location given is an existing city, the location will now be set to Trevignano", strconv.FormatInt(telegram.CHATID, 10))
		globals.CurrentLocation = "Trevignano"
	}

	dividedMessage := models.DivideMessages(result)

	if dividedMessage == "" {
		telegram.SendTelegramMessage("no raining today or tomorrow: "+weather.CurrentDateString+"    /    "+weather.TomorrowDateString+"\nLocation: "+globals.CurrentLocation, strconv.FormatInt(telegram.CHATID, 10))
	} else {
		telegram.SendTelegramMessage(dividedMessage+"\n\nLocation: "+globals.CurrentLocation, strconv.FormatInt(telegram.CHATID, 10))
	}

}
