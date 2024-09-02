package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemperatureDescription(t *testing.T) {
	tests := []struct {
		temp     int
		expected string
	}{
		{temp: 30, expected: "cold"},
		{temp: 50, expected: "moderate"},
		{temp: 90, expected: "hot"},
	}

	for _, tt := range tests {
		result := TemperatureDescription(tt.temp)
		if result != tt.expected {
			t.Errorf("temperatureDescription(%d) = %s; expected %s", tt.temp, result, tt.expected)
		}
	}
}

func TestGetWeatherData(t *testing.T) {
	// Mock response data
	mockResponse := `{
		"properties": {
			"periods": [{
				"name": "Today",
				"temperature": 72,
				"temperatureUnit": "F",
				"shortForecast": "Partly Cloudy"
			}]
		}
	}`

	// Create a new test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer ts.Close()

	// Replace the NWS base URL with the test server's URL
	oldNwsBaseURL := nwsBaseURL
	nwsBaseURL := ts.URL + "/"
	fmt.Println("nwsBaseURL", nwsBaseURL)
	defer func() { nwsBaseURL = oldNwsBaseURL }()

	// Call the function
	forecast, err := GetWeatherData("32", "81")
	if err != nil {
		t.Fatalf("getWeatherData() error: %v", err)
	}

	// Check the results
	if len(forecast.Properties.Periods) == 0 {
		t.Errorf("Expected at least one period in the forecast")
	} else {
		if forecast.Properties.Periods[0].ShortForecast != "Mostly Sunny" {
			t.Errorf("Expected 'Partly Cloudy', got '%s'", forecast.Properties.Periods[0].ShortForecast)
		}
	}
}
