package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JacklO0p/weather_forecast/api/location"
)

func GetWeatherFromLocation() (resp []map[string]string){
	latitude, longitude := location.GetCoordinates()
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + strconv.FormatFloat(latitude, 'f', 2, 64) + "&longitude=" + strconv.FormatFloat(longitude, 'f', 2, 64) + "&hourly=temperature_2m,relativehumidity_2m,precipitation,rain,showers,freezinglevel_height,surface_pressure,cloudcover,visibility,windspeed_80m,soil_temperature_6cm"

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error while getting weather values, %v", err)
		return []map[string]string{}
	}

	defer res.Body.Close()

	var meteo []map[string]string

	err = json.NewDecoder(res.Body).Decode(&meteo)
	if err != nil {
		fmt.Printf("Error while deconding, %v", err)
		return []map[string]string{}
	}

	fmt.Print("Latitude: ", latitude)
	fmt.Print("\nlongitude: ", longitude)
	fmt.Print("\n\n", meteo)

	return meteo
}
