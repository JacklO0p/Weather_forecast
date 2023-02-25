package telegram

import (
	"encoding/json"
	"fmt"

	"github.com/JacklO0p/weather_forecast/models"
)

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

	

}

func isItRaining(weather models.Weather) bool {
	
	for index, _ := range weather.Daily.PrecipitationHours {
		if weather.Daily.PrecipitationHours[index] != 0  || weather.Daily.RainSum[index] != 0 {
			return true
		}
	}

	return false
}

func userString(weather models.Weather) string {
	var message string = ""

	message += "RAIN ALERT!!\n\nWeather condition:\nT max {\n"
	for index, _ := range weather.Daily.ApparentTemperatureMax {
		message += "    " + weather[index].Daily.ApparentTemperatureMax
	}


	return message
}