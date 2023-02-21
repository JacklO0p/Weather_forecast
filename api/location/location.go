package location

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var location string = "Trevignano"

func GetLocation() string {
	return location
}

func GetCoordinates() (latitude float64, longitude float64, err error) {
	res, err := http.Get("https://nominatim.openstreetmap.org/search/Trevignano?format=json&addressdetails=1&limit=1")
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
		fmt.Printf("error while parsing[latitude], %v", err)
		return 0, 0, err
	}

	longitude, err = strconv.ParseFloat(location[0]["lon"].(string), 2)
	if err != nil {
		fmt.Printf("error while parsing[longitude]: %v", err)
		return 0, 0, err
	}

	return latitude, longitude, nil
}
