package controllers

import (
	"fmt"
	"net/http"

	"github.com/JacklO0p/weather_forecast/api/telegram"
	"github.com/JacklO0p/weather_forecast/api/weather"
	"github.com/labstack/echo/v4"
)

func GetWeather(c echo.Context) error {
	var result map[string]interface{}

	result = weather.GetWeatherFromLocation()
	if result == nil {
		fmt.Print("Error while getting the value, result is nil\n")

		telegram.SendTelegramMessage("Found error in the code,\n\nMake sure the location given is an existing city, the location will now be set to Trevignano", telegram.CHATID)
		telegram.CurrentLocation = "Trevignano"
		return c.JSON(http.StatusBadRequest, result)
	}

	dividedMessage := telegram.DivideMessages(result)

	if dividedMessage == "" {
		telegram.SendTelegramMessage("no raining today or tomorrow: " + weather.CurrentDateString + "    /    " + weather.TomorrowDateString + "\nLocation: " + telegram.CurrentLocation, telegram.CHATID)
	} else {
		telegram.SendTelegramMessage(dividedMessage + "\n\nLocation: " + telegram.CurrentLocation, telegram.CHATID)
	}

	return c.JSON(http.StatusOK, result)
}
