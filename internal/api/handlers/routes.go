package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)


func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(
		WeatherPath,
		GetWeatherData,
	).Methods(http.MethodGet)

	return router
}
