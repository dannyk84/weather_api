from fastapi import FastAPI
from fastapi.exceptions import HTTPException

from business import get_forecast
from integrators import WeatherAPIClient


app = FastAPI()

@app.get("/forecast")
def root(longitude: float, latitude: float):
    client = WeatherAPIClient()
    result = get_forecast(client=client, longitude=longitude, latitude=latitude)
    if result.is_err:
        raise HTTPException(status_code=400, detail="Error getting forecast.")
    
    forecast = result.unwrap()

    return {
        "short_forecast": forecast["short_forecast"],
        "temperature": forecast["temperature"],
        "characterization": forecast["characterization"],
    }
