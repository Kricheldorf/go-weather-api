package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

const CityBookmarksPath = "/city_bookmarks"
const CityBookmarkPath = "/city_bookmarks/{id}"
const WeatherPath = "/weather"

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(
		WeatherPath,
		GetWeatherData,
	).Methods(http.MethodGet)

	router.HandleFunc(
		CityBookmarksPath,
		ListCityBookmarks,
	).Methods(http.MethodGet)

	router.HandleFunc(
		CityBookmarkPath,
		GetCityBookmark,
	).Methods(http.MethodGet)

	router.HandleFunc(
		CityBookmarksPath,
		CreateCityBookmark,
	).Methods(http.MethodPost)

	router.HandleFunc(
		CityBookmarkPath,
		DeleteCityBookmark,
	).Methods(http.MethodDelete)

	return router
}
