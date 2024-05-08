package openweathermapapi

import (
	"github.com/h2non/gock"
	"testing"
)

func (latLonResponse LatLonResponse) Equals(other LatLonResponse) bool {
	return latLonResponse.City == other.City && latLonResponse.State == other.State && latLonResponse.Country == other.Country && latLonResponse.Lat == other.Lat && latLonResponse.Lon == other.Lon
}

func (weatherData WeatherData) Equals(other WeatherData) bool {
	return weatherData.Name == other.Name &&
		weatherData.Main.Temp == other.Main.Temp &&
		weatherData.Main.FeelsLike == other.Main.FeelsLike &&
		weatherData.Main.TempMin == other.Main.TempMin &&
		weatherData.Main.TempMax == other.Main.TempMax &&
		weatherData.Main.Humidity == other.Main.Humidity &&
		weatherData.Wind.Speed == other.Wind.Speed &&
		weatherData.Wind.Deg == other.Wind.Deg &&
		weatherData.Base == other.Base
}

func TestGetLatLonData(t *testing.T) {
	t.Run("Should fetch lat and log data related to Joinville, SC, BR", func(t *testing.T) {
		defer gock.Off()

		gock.New(openWeatherMapApiUrl).
			Get("/geo/1.0/direct").
			Reply(200).
			BodyString(`[{"name":"Joinville","local_names":{"ru":"Жоинвили","sr":"Жоинвиле","ko":"조인빌리","ja":"ジョインヴィレ","mk":"Џоинвил","lt":"Žoinvilis","la":"Ioannevilla","he":"ז'וינווילי","ka":"ჟოინვილი","zh":"若茵维莱","os":"Жоинвили","pt":"Joinville","bg":"Жойнвили","lv":"Žoinvile"},"lat":-26.3044898,"lon":-48.8486726,"country":"BR","state":"Santa Catarina"}]`)

		city := "Joinville"
		state := "SC"
		country := "Brazil"

		expectedOutput := []LatLonResponse{
			{
				City:    "Joinville",
				State:   "Santa Catarina",
				Country: "BR",
				Lat:     -26.3044898,
				Lon:     -48.8486726,
			},
		}

		result, err := GetLatLonData(city, state, country)

		if len(result) != len(expectedOutput) && !result[0].Equals(expectedOutput[0]) {
			t.Errorf("Unexpected result. Got %v, want %v", result, expectedOutput)
		}

		if err != nil {
			t.Errorf("Unexpected error. Got %v, want %v", err, nil)
		}
	})

	t.Run("Should return empty slice if city is not found", func(t *testing.T) {
		defer gock.Off()

		gock.New(openWeatherMapApiUrl).
			Get("/geo/1.0/direct").
			Reply(200).
			BodyString(`[]`)

		city := "not_a_city"
		state := "SC"
		country := "Brazil"

		var expectedOutput []LatLonResponse

		result, _ := GetLatLonData(city, state, country)

		if len(result) != len(expectedOutput) && !result[0].Equals(expectedOutput[0]) {
			t.Errorf("Unexpected result. Got %v, want %v", result, expectedOutput)
		}
	})
}

func TestGetWeather(t *testing.T) {
	t.Run("Should fetch weather data related to Joinville, SC, BR", func(t *testing.T) {
		defer gock.Off()

		gock.New(openWeatherMapApiUrl).
			Get("/data/2.5/weather").
			Reply(200).
			BodyString(`{"coord":{"lon":-48.8487,"lat":-26.3045},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],"base":"stations","main":{"temp":26.04,"feels_like":26.04,"temp_min":24.96,"temp_max":28.4,"pressure":1015,"humidity":81},"visibility":10000,"wind":{"speed":1.03,"deg":260},"clouds":{"all":75},"dt":1715000090,"sys":{"type":2,"id":2091989,"country":"BR","sunrise":1714988530,"sunset":1715028105},"timezone":-10800,"id":3459712,"name":"Joinville","cod":200}`)

		lat := -26.3044898
		lon := -48.8486726

		response, err := GetWeather(lat, lon)

		expectedWeatherData := WeatherData{
			Weather: []struct {
				Main        string `json:"main"`
				Description string `json:"description"`
			}{
				{
					Main:        "Clouds",
					Description: "broken clouds",
				},
			},
			Base: "stations",
			Main: struct {
				Temp      float64 `json:"temp"`
				FeelsLike float64 `json:"feels_like"`
				TempMin   float64 `json:"temp_min"`
				TempMax   float64 `json:"temp_max"`
				Humidity  int     `json:"humidity"`
			}{
				Temp:      26.04,
				FeelsLike: 26.04,
				TempMin:   24.96,
				TempMax:   28.4,
				Humidity:  81,
			},
			Wind: struct {
				Speed float64 `json:"speed"`
				Deg   int     `json:"deg"`
			}{
				Speed: 1.03,
				Deg:   260,
			},
			Name: "Joinville",
		}

		if response.Equals(expectedWeatherData) {
			t.Errorf("Unexpected result. Got %v, want %v", response.Name, "Joinville")
		}

		if err != nil {
			t.Errorf("Unexpected error. Got %v, want %v", err, nil)
		}
	})
}
