package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type LocationQuery struct {
	Name        string
	Latitude    string
	Longitude   string
	CountryCode string
	Timezone    string
	Province    string
}

type Location struct {
	Name        string  `json:"name"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	CountryCode string  `json:"country_code"`
	Timezone    string  `json:"timezone"`
	Province    string  `json:"admin1"`
}

type LocationResultsResponse struct {
	Results []Location `json:"results"`
}

const locationUrl = "https://geocoding-api.open-meteo.com/v1/search"

func SearchLocation(name string) ([]Location, error) {
	endpoint := fmt.Sprintf("%s?name=%s", locationUrl, url.QueryEscape(name))
	r, err := http.Get(endpoint)

	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	defer r.Body.Close()

	var locationResponse LocationResultsResponse

	if err := json.NewDecoder(r.Body).Decode(&locationResponse); err != nil {
		return nil, err //errors.New("no results for city")
	}

	return locationResponse.Results, nil
}
