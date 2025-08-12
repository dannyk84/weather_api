package business

import (
	"api/web-service-gin/api/integrators"
	"errors"
	"fmt"
)

type BusinessLayer struct {
	wac integrators.WeatherAPIClient
}

func NewBusinessLayer(wac integrators.WeatherAPIClient) BusinessLayer {
	return BusinessLayer{wac: wac}
}

// Functions
func (bl BusinessLayer) GetForecast(
	longitude float32,
	latitude float32,
) (*GetForecastResult, error) {
	forecast, err := bl.wac.GetForecast(longitude, latitude)
	if err != nil {
		return nil, errors.New("error getting forecast")
	}

	shortForecast := forecast.ShortForecast
	temperature := fmt.Sprintf("%v %v", forecast.Temperature, forecast.TemperatureUnit)
	characterization := getCharacterization(forecast.Temperature)

	return &GetForecastResult{
		ShortForecast:    shortForecast,
		Temperature:      temperature,
		Characterization: characterization,
	}, nil
}

func getCharacterization(temperature int32) string {
	if temperature > 80 {
		return "hot"
	}
	if temperature > 70 {
		return "moderate"
	}

	return "cold"
}
