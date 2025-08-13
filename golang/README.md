## Running the application
```
pip install -r requirements.txt

cd ./golang
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
To save some time, I didn't write any Golang tests, however you can take a look at my Python test file to see what scenarios I would cover. If you'd like, we can discuss how I would approach testing in Golang.
