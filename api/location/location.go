package location

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var location string = "trevignano"

func GetLocation() string {
	return location
}

func GetCoordinates() (latitude float64, longitude float64) {
	res, err := http.Get("https://nominatim.openstreetmap.org/search/Trevignano?format=json&addressdetails=1&limit=1")
	if err != nil {
		fmt.Printf("Error while getting coordinates, %v", err)
	}

	defer res.Body.Close()

	location := make(map[string]interface{})

	fmt.Print(location)

	err = json.NewDecoder(res.Body).Decode(&location)

	if err != nil {
		fmt.Printf("Error while decoding, %v", err)
		return 0, 0
	}

	latitude = location["lat"].(float64)
	longitude = location["lon"].(float64)

	return latitude, longitude
}
