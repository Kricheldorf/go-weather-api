package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-weather-api/internal/api/models"
	"net/http"
)

func CreateCityBookmark(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var cityBookmark models.CityBookmark
	err := json.NewDecoder(request.Body).Decode(&cityBookmark)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(responseWriter).Encode(map[string]string{"error": "Invalid request payload"})
		if err != nil {
			return
		}
		return
	}
	err = models.CreateCityBookmark(cityBookmark)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(responseWriter).Encode(map[string]string{"error": "Couldn't create city bookmark"})
		if err != nil {
			return
		}
		return
	}
	responseWriter.WriteHeader(http.StatusCreated)
	return
}

func ListCityBookmarks(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	cityBookmarks, err := models.GetCityBookmarks()
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(responseWriter).Encode(map[string]string{"error": "Failed to get city bookmarks"})
		if err != nil {
			return
		}
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	err = json.NewEncoder(responseWriter).Encode(cityBookmarks)
	if err != nil {
		return
	}
}

func GetCityBookmark(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	id := vars["id"]
	cityBookmark, err := models.GetCityBookmark(id)
	if err != nil {
		if cityBookmark.Id == "" {
			responseWriter.WriteHeader(http.StatusNotFound)
			err := json.NewEncoder(responseWriter).Encode(map[string]string{"error": "City bookmark not found"})
			if err != nil {
				return
			}
			return
		}

		responseWriter.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(responseWriter).Encode(map[string]string{"error": "Failed to get city bookmark"})
		if err != nil {
			return
		}
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	err = json.NewEncoder(responseWriter).Encode(cityBookmark)
	if err != nil {
		return
	}
}

func DeleteCityBookmark(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	id := vars["id"]

	err := models.DeleteCityBookmark(id)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(responseWriter).Encode(map[string]string{"error": "Failed to delete city bookmark"})
		if err != nil {
			return
		}
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	return
}
