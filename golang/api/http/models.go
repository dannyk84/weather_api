package http

type GetForecastRequest struct {
	Longitude float32 `form:"longitude" binding:"required"`
	Latitude  float32 `form:"latitude" binding:"required"`
}

type GetForecastResponse struct {
	ShortForecast    string `json:"short_forecast"`
	Temperature      string `json:"temperature"`
	Characterization string `json:"characterization"`
}
