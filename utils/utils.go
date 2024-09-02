package utils

import (
	"encoding/json"
	"fmt"
	"github.com/rajumscis16/weather-service/dto"
	"net/http"
)

// TODO we can read this api from config in the feature
const nwsBaseURL = "https://api.weather.gov/gridpoints/TOP/"

// TemperatureDescription find characterization of whether the temperature
func TemperatureDescription(temp int) string {
	switch {
	case temp <= 40:
		return "cold"
	case temp >= 85:
		return "hot"
	default:
		return "moderate"
	}
}

// GetWeatherData connect to the NSW API and get details
func GetWeatherData(lat, lon string) (*dto.Forecast, error) {
	// Build the URL to get the forecast
	url := fmt.Sprintf("%s%s,%s/forecast", nwsBaseURL, lat, lon)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var forecast dto.Forecast
	err = json.NewDecoder(resp.Body).Decode(&forecast)
	if err != nil {
		return nil, err
	}

	return &forecast, nil
}
