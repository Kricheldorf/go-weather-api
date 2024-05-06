package handlers

import (
	"encoding/json"
	openweathermapapi "go-weather-api/internal/api/services"
	"net/http"
)

func GetWeatherData(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	city := request.URL.Query().Get("city")
	state := request.URL.Query().Get("state")
	country := request.URL.Query().Get("country")

	latLonData, err := openweathermapapi.GetLatLonData(city, state, country)

	if err != nil || len(latLonData) == 0 {
		couldNotGetLatLonData := map[string]string{"error": "Failed to get latitude and longitude data"}
		responseWriter.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(responseWriter).Encode(couldNotGetLatLonData)
		if err != nil {
			return
		}
		return
	}

	weatherInformation, err := openweathermapapi.GetWeather(latLonData[0].Lat, latLonData[0].Lon)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(responseWriter).Encode(map[string]string{"error": "Failed to get weather data"})
		if err != nil {
			return
		}
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	err = json.NewEncoder(responseWriter).Encode(weatherInformation)
	if err != nil {
		return
	}
}
