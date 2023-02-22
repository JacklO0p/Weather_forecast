package telegram

import (
	"encoding/json"
	"fmt"

	"github.com/JacklO0p/weather_forecast/models"
)

func DivideMessages(Message map[string]interface{}) string {
	stringMessage := GetString(Message)

	return stringMessage
}

func GetString(Message map[string]interface{}) string {
	jsonBytes, err := json.Marshal(Message)
	if err != nil {
		fmt.Println("Error while marshaling data to JSON:", err)
		return ""
	}

	weatherData := models.Weather{}

	err = json.Unmarshal(jsonBytes, &weatherData)
	if err != nil {
		fmt.Print("\nError while unmarshaling, ", err)
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
			return true
		}
	}

	return false
}

func GetFormattedString(Message models.Weather) (res string) {

	res += "***Rain Alert!***\n\n"
	
	res += "Apparent Temperature {\n      max: "
	
	for index, _ := range Message.Daily.ApparentTemperatureMax {
		res += fmt.Sprintf("%v", Message.Daily.ApparentTemperatureMax[index]) + ";   "
	}

	res += "\n      min: "
	
	for index, _ := range Message.Daily.ApparentTemperatureMin {
		res += fmt.Sprintf("%v", Message.Daily.ApparentTemperatureMin[index]) + ";   "
	}
	
	res += "\n}\n\nRaining hours {\n      "
	
	for index, _ := range Message.Daily.PrecipitationHours {
		res += fmt.Sprintf("%v", Message.Daily.PrecipitationHours[index]) + ";  "
	}
	
	res += "\n}\n\nTotal rain [mm] {\n      "
	
	for index, _ := range Message.Daily.RainSum {
		res += fmt.Sprintf("%v", Message.Daily.RainSum[index]) + ";  "
	}

	res += "\n}"


	return res
}
