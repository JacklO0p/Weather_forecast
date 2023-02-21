package models

type Weather struct {
	CurrentWeather struct {
		Temperature   float64 `json:"temperature"`
		Time          string  `json:"time"`
		Weathercode   int     `json:"weathercode"`
		Winddirection int     `json:"winddirection"`
		Windspeed     float64 `json:"windspeed"`
	} `json:"current_weather"`
	Daily struct {
		ApparentTemperatureMax []float64 `json:"apparent_temperature_max"`
		ApparentTemperatureMin []float64 `json:"apparent_temperature_min"`
		PrecipitationHours     []int     `json:"precipitation_hours"`
		RainSum                []int     `json:"rain_sum"`
		Temperature2MMax       []float64 `json:"temperature_2m_max"`
		Temperature2MMin       []float64 `json:"temperature_2m_min"`
		Time                   []string  `json:"time"`
	} `json:"daily"`
	DailyUnits struct {
		ApparentTemperatureMax string `json:"apparent_temperature_max"`
		ApparentTemperatureMin string `json:"apparent_temperature_min"`
		PrecipitationHours     string `json:"precipitation_hours"`
		RainSum                string `json:"rain_sum"`
		Temperature2MMax       string `json:"temperature_2m_max"`
		Temperature2MMin       string `json:"temperature_2m_min"`
		Time                   string `json:"time"`
	} `json:"daily_units"`
	Elevation            int     `json:"elevation"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
}