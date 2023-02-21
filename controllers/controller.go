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
		fmt.Print("Error while getting the value, result is nil")

		telegram.SendTelegramMessage("Found error in the code, the program will now shut down", telegram.CHATID)
		return c.JSON(http.StatusBadRequest, "")
	}

	fmt.Print(result, "\n")
	dividedMessage := telegram.DivideMessages(result)

	if dividedMessage == "" {
		telegram.SendTelegramMessage("no raining today: " + weather.CurrentDateString, telegram.CHATID)
	} else {
		telegram.SendTelegramMessage(dividedMessage, telegram.CHATID)
	}

	return c.JSON(http.StatusOK, result)
}
