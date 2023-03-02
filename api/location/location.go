package location

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/JacklO0p/weather_forecast/globals"
	"github.com/JacklO0p/weather_forecast/models"
)

func GetCoordinates(loca string) (latitude float64, longitude float64, err error) {
	var res *http.Response

	if isValid(loca) {
		res, err = http.Get("https://nominatim.openstreetmap.org/search/" + url.QueryEscape(loca) + "?format=json&addressdetails=1&limit=1")
	} else {
		return -100000000, -100000000, nil
	}

	if err != nil {
		fmt.Printf("Error while getting coordinates, %v", err)
	}

	defer res.Body.Close()

	var location []map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&location)

	if err != nil {
		fmt.Printf("Error while decoding, %v", err)
		return 0, 0, nil
	}

	latitude, err = strconv.ParseFloat(location[0]["lat"].(string), 2)
	if err != nil {
		fmt.Printf("Error while parsing[latitude], %v", err)
		return 0, 0, err
	}

	longitude, err = strconv.ParseFloat(location[0]["lon"].(string), 2)
	if err != nil {
		fmt.Printf("Error while parsing[longitude]: %v", err)
		return 0, 0, err
	}

	user := models.User{
		Location: loca,
	}

	globals.Db.Where("Location=?", loca).First(&user)

	user.Latitude = strconv.FormatFloat(latitude, 'f', 2, 64)
	user.Longitude = strconv.FormatFloat(longitude, 'f', 2, 64)

	models.UpdateUser(&user)

	return latitude, longitude, nil
}

func isValid(city string) bool {

	res, err := http.Get("http://www.weather-forecast.com/locations/ac_location_name?query=" + url.QueryEscape(city))
	if err != nil {
		fmt.Print("\nError while getting the city api ", err, "\n")
	}

	cityCheck := models.CityValidator{}

	err = json.NewDecoder(res.Body).Decode(&cityCheck)

	if cityCheck.NearObjCount != 0 {
		return true
	}

	return false
}
