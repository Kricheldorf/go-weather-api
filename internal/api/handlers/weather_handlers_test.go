package handlers

import (
	"github.com/gorilla/mux"
	"github.com/h2non/gock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWeatherDataHandler(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.openweathermap.org").
		Get("/data/2.5/weather").
		Reply(200).
		BodyString(`{"coord":{"lon":-48.8487,"lat":-26.3045},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],"base":"stations","main":{"temp":26.04,"feels_like":26.04,"temp_min":24.96,"temp_max":28.4,"pressure":1015,"humidity":81},"visibility":10000,"wind":{"speed":1.03,"deg":260},"clouds":{"all":75},"dt":1715000090,"sys":{"type":2,"id":2091989,"country":"BR","sunrise":1714988530,"sunset":1715028105},"timezone":-10800,"id":3459712,"name":"Joinville","cod":200}`)

	gock.New("https://api.openweathermap.org").
		Get("/geo/1.0/direct").
		Reply(200).
		BodyString(`[{"name":"Joinville","local_names":{"ru":"Жоинвили","sr":"Жоинвиле","ko":"조인빌리","ja":"ジョインヴィレ","mk":"Џоинвил","lt":"Žoinvilis","la":"Ioannevilla","he":"ז'וינווילי","ka":"ჟოინვილი","zh":"若茵维莱","os":"Жоинвили","pt":"Joinville","bg":"Жойнвили","lv":"Žoinvile"},"lat":-26.3044898,"lon":-48.8486726,"country":"BR","state":"Santa Catarina"}]`)

	req, err := http.NewRequest("GET", "/weather?city=joinville&city=joinville&state=sc&country=brazil", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/weather", GetWeatherData).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "{\"weather\":[{\"main\":\"Clouds\",\"description\":\"broken clouds\"}],\"base\":\"stations\",\"main\":{\"temp\":26.04,\"feels_like\":26.04,\"temp_min\":24.96,\"temp_max\":28.4,\"humidity\":81},\"wind\":{\"speed\":1.03,\"deg\":260,\"gust\":0},\"name\":\"Joinville\"}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
