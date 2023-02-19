package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JacklO0p/weather_forecast/api/location"
)

func GetWeatherFromLocation() (resp map[string]interface{}) {
	latitude, longitude, err := location.GetCoordinates()
	if err != nil {

	}

	
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + strconv.FormatFloat(latitude, 'f', 2, 64) + "&longitude=" + strconv.FormatFloat(longitude, 'f', 2, 64) + "&daily=temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,rain_sum,precipitation_hours&current_weather=true&timezone=Europe%2FBerlin&start_date=2023-02-19&end_date=2023-02-26"
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error while getting weather values, %v", err)
		return map[string]interface{}{}
	}

	defer res.Body.Close()
	var meteo map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&meteo)
	if err != nil {
		fmt.Printf("Error while deconding, %v", err)
		return map[string]interface{}{}
	}

	return meteo
}
