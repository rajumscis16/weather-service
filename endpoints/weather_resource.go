package endpoints

import (
	"encoding/json"
	"github.com/rajumscis16/weather-service/utils"
	"net/http"
	"strconv"
)

type WeatherResourceImp struct{}

type weatherResource interface {
	WeatherHandler(w http.ResponseWriter, r *http.Request)
}

// WeatherHandler find short forecast and characterization of whether
func (weatherResource WeatherResourceImp) WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Get latitude and longitude from query parameters
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" || lon == "" {
		http.Error(w, "Please provide latitude and longitude", http.StatusBadRequest)
		return
	}

	// Get the weather data from the NWS API
	forecast, err := utils.GetWeatherData(lat, lon)
	if err != nil {
		http.Error(w, "Failed to get weather data", http.StatusInternalServerError)
		return
	}

	// Get today's forecast
	todayForecast := forecast.Properties.Periods[0]

	// Determine temperature characterization
	tempDesc := utils.TemperatureDescription(todayForecast.Temperature)

	// Prepare the response
	response := map[string]string{
		"short_forecast":   todayForecast.ShortForecast,
		"temperature":      strconv.Itoa(todayForecast.Temperature) + "°" + todayForecast.TemperatureUnit,
		"characterization": tempDesc,
	}

	// Encode and send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
