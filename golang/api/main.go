package main

import (
	"api/web-service-gin/api/business"
	"api/web-service-gin/api/http"
	"api/web-service-gin/api/integrators"

	"github.com/gin-gonic/gin"
)

func main() {
	wac := integrators.NewWeatherAPICient()
	bl := business.NewBusinessLayer(wac)
	hl := http.NewHttpLayer(bl)

	router := gin.Default()
	router.GET("/forecast", hl.GetForecast)

	router.Run()
}
