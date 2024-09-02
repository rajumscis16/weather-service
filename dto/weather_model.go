package dto

// Forecast NWS model
type Forecast struct {
	Properties struct {
		Periods []struct {
			Name            string `json:"name"`
			Temperature     int    `json:"temperature"`
			TemperatureUnit string `json:"temperatureUnit"`
			ShortForecast   string `json:"shortForecast"`
		} `json:"periods"`
	} `json:"properties"`
}
