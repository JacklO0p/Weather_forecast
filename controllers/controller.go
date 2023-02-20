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
		return c.JSON(http.StatusBadRequest, "")
	}

	fmt.Print(result, "\n")
	dividedMessage := telegram.DivideMessages(result)

	var r string = ""

	for index := range dividedMessage {
		r += dividedMessage[index]
	}
	//finalMessage := telegram.FixMessage(dividedMessage)

	//fmt.Print(finalMessage)

	fmt.Print(r)
	telegram.SendTelegramMessage(r, telegram.CHATID)

	return c.JSON(http.StatusOK, result)
}
