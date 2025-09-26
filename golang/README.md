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
