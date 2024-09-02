Use of this API

1.    Accepts latitude and longitude coordinates

2.    Returns the short forecast for that area for Today (“Partly Cloudy” etc)

3.    Returns a characterization of whether the temperature is “hot”, “cold”, or “moderate” (use your discretion on mapping temperatures to each type)

4.    Use the National Weather Service API Web Service as a data source.

How to run the service

`go build`

`go run main.go`

How to run tests 

`go test -v ./...`

How to use the API

`GET  http://localhost:8080/weather?lat=32&lon=81
`

And the resposne is

`{
"characterization": "moderate",
"short_forecast": "Mostly Sunny",
"temperature": "78°F"
}`

Based on 

[
https://www.weather.gov/documentation/services-web-api#/
](https://www.weather.gov/documentation/services-web-api#/)

