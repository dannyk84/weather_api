import logging
import requests

from option import Err, Ok, Result


class WeatherAPIClient:
    BASE_URL = "https://api.weather.gov"

    def __init__(self):
        logging.info("Starting up WeatherAPIClient")

    def get_forecast(self, longitude: float, latitude: float) -> Result[dict, str]:
        """
        Returns forecast information for this afternoon. We're first calling the
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
        """

        points_url = f"{self.BASE_URL}/points/{longitude},{latitude}"
        points_response = requests.get(points_url).json()
        err = forecast_response.get("detail")
        if err:
            err = f"Error calling points endpoint | err={err}"
            logging.error(err)
            return Err(err)

        forecast_url = points_response["properties"]["forecast"]
        forecast_response = requests.get(forecast_url).json()
        err = forecast_response.get("detail")
        if err:
            err = f"Error calling forecast endpoint | err={err}"
            logging.error(err)
            return Err(err)

        # This gets the forecast for this afternoon
        forecast_values = forecast_response["properties"]["periods"][0]

        return Ok(
            {
                "temperature": forecast_values["temperature"],
                "temperature_unit": forecast_values["temperatureUnit"],
                "short_forecast": forecast_values["shortForecast"],
            }
        )
