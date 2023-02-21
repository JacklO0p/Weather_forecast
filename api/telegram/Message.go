package telegram

import (
	"encoding/json"
	"fmt"

	"github.com/JacklO0p/weather_forecast/models"
)

func DivideMessages(message map[string]interface{}) string {
	stringMessage := GetString(message)

	fmt.Print("\n\n\n", stringMessage, "\n\n\n\n")

	return stringMessage
}

func GetString(message map[string]interface{}) string {
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("error while marshaling data to JSON:", err)
		return ""
	}

	weatherData := models.Weather{}

	err = json.Unmarshal(jsonBytes, &weatherData)
	if err != nil {
		fmt.Print("\nerror while unmarshaling, ", err)
	}

	var checkRain = GetRains(weatherData)

	if checkRain {
		return GetFormattedString(weatherData)
	}

	return ""
}

func GetRains(weather models.Weather) (bool) {

	for index, _ := range weather.Daily.RainSum {
		if weather.Daily.RainSum[index] != 0 {
			fmt.Print("\n\n\n\n", weather.Daily.RainSum[index], "\n\n\n\n")
			return true
		}
	}

	return false
}

func GetFormattedString(message models.Weather) (res string) {

	res += "***Rain Alert!***\n\n"
	
	res += "Apparent Temperature {\n      max: "
	
	for index, _ := range message.Daily.ApparentTemperatureMax {
		res += fmt.Sprintf("%v", message.Daily.ApparentTemperatureMax[index]) + ";   "
	}

	res += "\n      min: "
	
	for index, _ := range message.Daily.ApparentTemperatureMin {
		res += fmt.Sprintf("%v", message.Daily.ApparentTemperatureMin[index]) + ";   "
	}
	
	res += "\n}\n\nRaining hours {\n      "
	
	for index, _ := range message.Daily.PrecipitationHours {
		res += fmt.Sprintf("%v", message.Daily.PrecipitationHours[index]) + ";  "
	}
	
	res += "\n}\n\nTotal rain [mm] {\n      "
	
	for index, _ := range message.Daily.RainSum {
		res += fmt.Sprintf("%v", message.Daily.RainSum[index]) + ";  "
	}

	res += "\n}"


	return res
}
