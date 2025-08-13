package integrators

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type WeatherAPIClient struct {
	url string
}

func NewWeatherAPICient() WeatherAPIClient {
	log.Print("Starting up WeatherAPIClient")

	return WeatherAPIClient{url: "https://api.weather.gov"}
}

// Endpoints
func (wac WeatherAPIClient) GetForecast(
	longitude float32,
	latitude float32,
) (*Forecast, error) {
	/*
		Returns the most recent forecast data for today. We're first calling the
		'points' endpoint, which returns a url for the 'forecast' endpoint. We then
		call the 'forecast' endpoint to get the data we actually need.

		Points
		####################
		Documentation: https://www.weather.gov/documentation/services-web-api#/default/point
		Example:       https://api.weather.gov/points/39.7456,-97.0892

		Forecast
		####################
		Documentation: https://www.weather.gov/documentation/services-web-api#/default/gridpoint_forecast
		Example:       https://api.weather.gov/gridpoints/TOP/32,81/forecast
	*/

	// Points endpoint
	pointsUrl := fmt.Sprintf("%v/points/%v,%v", wac.url, longitude, latitude)
	pointsRespData, err := sendRequest(pointsUrl)
	if err != nil {
		log.Printf("error calling points endpoint | err=%v", err)
		return nil, err
	}

	var pointsObj PointsResponse
	json.Unmarshal(pointsRespData, &pointsObj)
	if pointsObj.Error != nil {
		errMsg := *pointsObj.Error
		log.Printf("points response has an error | err=%v", errMsg)
		return nil, errors.New(errMsg)
	}

	// Forecast endpoint
	forecastUrl := pointsObj.Properties.Forecast
	forecastRespData, err := sendRequest(forecastUrl)
	if err != nil {
		log.Printf("error calling forecast endpoint | err=%v", err)
		return nil, err
	}

	var forecastObj ForecastResponse
	json.Unmarshal(forecastRespData, &forecastObj)
	if forecastObj.Error != nil {
		errMsg := *forecastObj.Error
		log.Printf("forecast response has an error | err=%v", errMsg)
		return nil, errors.New(errMsg)
	}

	// This gets the most recent forecast for today
	forecastValues := forecastObj.Properties.Periods[0]

	return &Forecast{
		Temperature:     forecastValues.Temperature,
		TemperatureUnit: forecastValues.TemperatureUnit,
		ShortForecast:   forecastValues.ShortForecast,
	}, nil
}
