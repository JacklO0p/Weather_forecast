package telegram

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/JacklO0p/weather_forecast/api/weather"
	"github.com/JacklO0p/weather_forecast/globals"
	"github.com/JacklO0p/weather_forecast/models"
)

func GetReport(loca string) string {

	result := weather.GetWeatherFromLocation(loca)
	if result == nil {
		return ""
	}

	dividedMessage := DivideMessage(result)
	if dividedMessage == "" {
		return "No raining tomorrow\n"
	}

	return dividedMessage
}

func DivideMessage(mes map[string]interface{}) string {
	jsonByte, err := json.Marshal(mes)
	if err != nil {
		fmt.Print("Error while marshalling message", err)
		return ""
	}

	weatherData := models.Weather{}

	json.Unmarshal(jsonByte, &weatherData)

	checkIfRain := isItRaining(weatherData)

	if checkIfRain {
		return userString(weatherData)
	}

	return ""
}

func isItRaining(weather models.Weather) bool {

	for index := range weather.Daily.PrecipitationHours {
		if weather.Daily.PrecipitationHours[index] != 0 || weather.Daily.RainSum[index] != 0 {
			return true
		}
	}

	return false
}

func userString(weather models.Weather) string {
	var message string = ""

	message += "RAIN ALERT!!\n\nWeather condition [\n\n    Temp max:  "
	for index := range weather.Daily.ApparentTemperatureMax {
		message += "    " + strconv.FormatFloat(weather.Daily.ApparentTemperatureMax[index], 'f', 2, 64)
	}

	message += "\n\n    Temp min:  "

	for index := range weather.Daily.ApparentTemperatureMax {
		message += "    " + strconv.FormatFloat(weather.Daily.ApparentTemperatureMin[index], 'f', 2, 64)
	}

	message += "\n\n    Hours of rain:  "

	for index := range weather.Daily.PrecipitationHours {
		message += "    " + strconv.Itoa(weather.Daily.PrecipitationHours[index])
	}

	message += "\n\n    Total rain:  "

	for index := range weather.Daily.RainSum {
		message += "    " + strconv.FormatFloat(weather.Daily.RainSum[index], 'f', 2, 64)
	}

	message += "\n]"

	message += "\n\nDate:  " + globals.CurrentDateString + "  /  " + globals.TomorrowDateString
	
	return message
}
