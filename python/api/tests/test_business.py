from option import Err, Ok
from unittest.mock import patch

from api.business import get_characterization, get_forecast
from api.integrators import WeatherAPIClient


@patch("api.business.get_characterization")
@patch("api.integrators.WeatherAPIClient.get_forecast")
class TestGetForecast:
    def test_successfully_get_forecast(
        self,
        mock_get_forecast,
        mock_get_characterization,
    ):
        mock_get_forecast.return_value = Ok(
            {
                "temperature": 90,
                "temperature_unit": "F",
                "short_forecast": "Blazing",
            }
        )
        mock_get_characterization.return_value = "hot"

        client = WeatherAPIClient()
        longitude = 39.7456
        latitude = -97.0892

        # Assertions
        result = get_forecast(client=client, longitude=longitude, latitude=latitude)
        assert result.is_ok

        actual = result.unwrap()
        assert actual["short_forecast"] == "Blazing"
        assert actual["temperature"] == "90 F"
        assert actual["characterization"] == "hot"

        mock_get_forecast.assert_called_once_with(
            longitude=longitude,
            latitude=latitude,
        )
        mock_get_characterization.assert_called_once_with(temperature=90)

    def test_api_error(
        self,
        mock_get_forecast,
        mock_get_characterization,
    ):
        mock_get_forecast.return_value = Err("Error")

        client = WeatherAPIClient()
        longitude = 39.7456
        latitude = -97.0892

        # Assertions
        result = get_forecast(client=client, longitude=longitude, latitude=latitude)
        assert result.is_err
        assert result.unwrap_err() == "Error getting forecast"

        mock_get_forecast.assert_called_once_with(
            longitude=longitude,
            latitude=latitude,
        )
        assert mock_get_characterization.call_count == 0


class TestGetCharacterization:
    def test_hot(self):
        actual = get_characterization(81)
        assert actual == "hot"

    def test_moderate(self):
        # Tests max boundary
        actual = get_characterization(80)
        assert actual == "moderate"

        # Tests min boundary
        actual = get_characterization(71)
        assert actual == "moderate"

    def test_cold(self):
        actual = get_characterization(70)
        assert actual == "cold"
