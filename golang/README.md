## Running the application
```
cd ./api
go run main.go
```

## Calling the endpoint
```
# Request
curl "http://127.0.0.1:8080/forecast?longitude=39.7456&latitude=-97.0892"

# Response
{"short_forecast":"Sunny","temperature":"85 F","characterization":"hot"}
```

## Running tests
- Due to time, I didn't write any Golang tests, however you can take a look at some Python tests I've written to see what scenarios I would cover.
- If you'd like, we can also discuss how I would approach testing in Golang.

## Some improvements I can make
- Init a `context` in `main.go` and pass it to the other layers.
- Move some of the marshaling logic in `WeatherAPIClient.GetForecast()` to a reusable function.
- Break out the multiple API calls in `WeatherAPIClient.GetForecast()` into two separate functions (`GetForecast` and `GetPoints`)
