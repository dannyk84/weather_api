package integrators

type Forecast struct {
	Temperature     int32
	TemperatureUnit string
	ShortForecast   string
}

// Endpoints

// Points
type PointsResponse struct {
	Properties *PointsProperties `json:"properties"`
	Error      *string           `json:"detail"`
}

type PointsProperties struct {
	Forecast string `json:"forecast"`
}

// Forecast
type ForecastResponse struct {
	Properties *ForecastProperties `json:"properties"`
	Error      *string             `json:"detail"`
}

type ForecastProperties struct {
	Periods []ForecastPeriods `json:"periods"`
}

type ForecastPeriods struct {
	Temperature     int32  `json:"temperature"`
	TemperatureUnit string `json:"temperatureUnit"`
	ShortForecast   string `json:"shortForecast"`
}
