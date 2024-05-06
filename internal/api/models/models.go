package models

import (
	"database/sql"
)

var DB *sql.DB

type CityBookmark struct {
	Id        string `json:"id" db:"id"`
	City      string `json:"city" db:"city"`
	State     string `json:"state" db:"state"`
	Country   string `json:"country" db:"country"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

func CreateCityBookmark(cityBookmark CityBookmark) error {
	_, err := DB.Exec("INSERT INTO city_bookmarks (city, state, country) VALUES ($1, $2, $3)", cityBookmark.City, cityBookmark.State, cityBookmark.Country)
	return err
}

func GetCityBookmarks() ([]CityBookmark, error) {
	rows, err := DB.Query("SELECT * FROM city_bookmarks")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var cityBookmarks []CityBookmark
	for rows.Next() {
		var cityBookmark CityBookmark
		err := rows.Scan(&cityBookmark.Id, &cityBookmark.City, &cityBookmark.State, &cityBookmark.Country, &cityBookmark.CreatedAt)
		if err != nil {
			return nil, err
		}
		cityBookmarks = append(cityBookmarks, cityBookmark)
	}

	return cityBookmarks, nil
}

func DeleteCityBookmark(id string) error {
	_, err := DB.Exec("DELETE FROM city_bookmarks WHERE id = $1", id)
	return err
}

func GetCityBookmark(id string) (CityBookmark, error) {
	var cityBookmark CityBookmark
	err := DB.QueryRow("SELECT * FROM city_bookmarks WHERE id = $1", id).Scan(&cityBookmark.Id, &cityBookmark.City, &cityBookmark.State, &cityBookmark.Country, &cityBookmark.CreatedAt)
	if err != nil {
		return CityBookmark{}, err
	}

	return cityBookmark, nil
}
