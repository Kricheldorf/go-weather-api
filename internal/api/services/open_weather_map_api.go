package openweathermapapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherData struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Name string `json:"name"`
}

type LatLonResponse struct {
	City    string  `json:"name"`
	State   string  `json:"state"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

const openWeatherMapApiUrl = "https://api.openweathermap.org"

func GetWeather(lat float64, lon float64) (WeatherData, error) {
	apiURL := fmt.Sprintf("%s/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s", openWeatherMapApiUrl, lat, lon, os.Getenv("OPEN_WEATHER_MAP_API_KEY"))

	response, err := http.Get(apiURL)
	if err != nil {
		return WeatherData{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	weatherData := WeatherData{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&weatherData)

	if err != nil {
		return WeatherData{}, err
	}

	return weatherData, nil
}

func GetLatLonData(city string, state string, country string) ([]LatLonResponse, error) {
	apiURL := fmt.Sprintf("%s/geo/1.0/direct?q=%s,%s,%s&appid=%s", openWeatherMapApiUrl, city, state, country, os.Getenv("OPEN_WEATHER_MAP_API_KEY"))

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return []LatLonResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	var citiesLatLon []LatLonResponse
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&citiesLatLon)

	if err != nil {
		return []LatLonResponse{}, err
	}

	return citiesLatLon, nil
}
