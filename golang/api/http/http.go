package http

import (
	"api/web-service-gin/api/business"

	"github.com/gin-gonic/gin"
)

type HttpLayer struct {
	bl business.BusinessLayer
}

func NewHttpLayer(bl business.BusinessLayer) HttpLayer {
	return HttpLayer{bl: bl}
}

// Endpoints
func (hl HttpLayer) GetForecast(c *gin.Context) {
	var requestData GetForecastRequest
	err := c.ShouldBindQuery(&requestData)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := hl.bl.GetForecast(requestData.Longitude, requestData.Latitude)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		200,
		GetForecastResponse{
			ShortForecast:    result.ShortForecast,
			Temperature:      result.Temperature,
			Characterization: result.Characterization,
		},
	)

}
