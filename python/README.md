## Running the application
```
pip install -r requirements.txt

cd ./api
uvicorn main:app --reload
```

## Calling the endpoint
```
# Request
curl "http://127.0.0.1:8000/forecast?longitude=39.7456&latitude=-97.0892"

# Response
{"short_forecast":"Sunny","temperature":"85 F","characterization":"hot"}
```

## Running tests
Typically I try to reach close to 100% coverage; however for this project, I only wrote business layer tests.
```
pytest
```
