from option import Result, Ok, Err

from integrators import WeatherAPIClient


def get_forecast(
    client: WeatherAPIClient,
    longitude: float,
    latitude: float,
) -> Result[dict, str]:
    result = client.get_forecast(longitude=longitude, latitude=latitude)
    if result.is_err:
        return Err("Error getting forecast")
    
    forecast = result.unwrap()

    short_forecast = forecast["short_forecast"]
    temperature = f"{forecast['temperature']} {forecast['temperature_unit']}"
    characterization = get_characterization(temperature=forecast["temperature"])

    return Ok(
        {
            "short_forecast": short_forecast,
            "temperature": temperature,
            "characterization": characterization,
        }
    )


def get_characterization(temperature: int) -> str:
    if temperature > 80:
        return "hot"
    if temperature > 70:
        return "moderate"
    
    return "cold"
