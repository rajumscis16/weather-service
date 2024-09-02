package cmd

import (
	"github.com/rajumscis16/weather-service/endpoints"
	"log"
	"net/http"
)

var (
	Application application = &app{}
)

type app struct{}

type application interface {
	Run() error
}

// Run start the service
func (app *app) Run() error {
	log.Println(" executing application")
	var wr = endpoints.WeatherResourceImp{}
	http.HandleFunc("/weather", wr.WeatherHandler)
	log.Println("Starting server on :8080...")

	return http.ListenAndServe(":8080", nil)
}
