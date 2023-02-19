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

	finalRes := fmt.Sprintf("%v", result)
	fmt.Print(finalRes)
	telegram.SendTelegramMessage(finalRes, telegram.CHATID)

	return c.JSON(http.StatusOK, finalRes)
}
