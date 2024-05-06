# Go Weather API

This app is a simple weather API that fetches weather data from OpenWeatherMap API and returns to the user.
Other functionalities include CRUD operations for bookmarked cities (except the Update part).

This is my first Go project, so I'm still learning the language and its best practices. I created this project for learning purposes.

## Example requests

### Fetching weather data

```bash
$ curl --request GET \
  --url 'http://localhost:8080/weather?city=joinville%2Csc%2Cbrazil&state=sc&country=brazil'
```

### Bookmarking a city

```bash
$ curl --request POST \
  --url 'http://localhost:8080/city_bookmarks' \
  --header 'Content-Type: application/json' \
  --data '{
	"city": "joinville",
	"state": "sc",
	"country": "brazil"
}'
```

### Fetching bookmarked cities

```bash
$ curl --request GET \
  --url 'http://localhost:8080/city_bookmarks'
```

### Deleting a bookmarked city

```bash
$ curl --request DELETE \
  --url 'http://localhost:8080/city_bookmarks/7'
```

### Fetching a bookmarked city

```bash
$ curl --request GET \
  --url 'http://localhost:8080/city_bookmarks/7'
```

## Setting up database

```bash
$ docker-compose up -d
$ migrate -path=internal/database/migrations -database "postgresql://go_weather:go_weather@localhost:5432/go_weather?sslmode=disable" -verbose up
```

## Configuring the server

```bash
$ cp .env.sample .env
```

## Running the server

```bash
$ go run cmd/app/main.go
```
