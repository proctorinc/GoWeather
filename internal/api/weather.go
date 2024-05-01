package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/go-querystring/query"
)

type DailyTempuratures struct {
	Date           []string  `json:"time"`
	MinTempurature []float32 `json:"temperature_2m_min"`
	MaxTempurature []float32 `json:"temperature_2m_max"`
}

type DailyForecastResponse struct {
	Latitude  float32           `json:"latitude"`
	Longitude float32           `json:"longitude"`
	Timezone  string            `json:"timezone"`
	Daily     DailyTempuratures `json:"daily"`
	Reason    string            `json:"reason"`
	Error     bool              `json:"error"`
}

type DailyForecastData = string

const (
	MaxTempurature              DailyForecastData  = "temperature_2m_max"
	MinTempurature              DailyForecastData  = "temperature_2m_min"
	RainSum                     HourlyForecastData = "rain_sum"
	ShowersSum                  HourlyForecastData = "showers_sum"
	SnowfallSum                 HourlyForecastData = "snowfall_sum"
	PrecipitationHours          HourlyForecastData = "precipitation_hours"
	PrecipitationProbabilityMax HourlyForecastData = "precipitation_probability_max"
	Sunrise                     HourlyForecastData = "sunrise"
	Sunset                      HourlyForecastData = "sunset"
)

type HourlyTempuratures struct {
	Time        []string  `json:"time"`
	Tempurature []float32 `json:"temperature_2m"`
}

type HourlyForecastResponse struct {
	Latitude  float32            `json:"latitude"`
	Longitude float32            `json:"longitude"`
	Timezone  string             `json:"timezone"`
	Hourly    HourlyTempuratures `json:"hourly"`
	Reason    string             `json:"reason"`
	Error     bool               `json:"error"`
}

type HourlyForecastData = string

const (
	Tempurature              HourlyForecastData = "temperature_2m"
	ApparentTempurature      HourlyForecastData = "apparent_temperature"
	RelativeHumidity         HourlyForecastData = "relative_humidity_2m"
	WeatherCode              HourlyForecastData = "weather_code"
	Precipitation            HourlyForecastData = "precipitation"
	PrecipitationProbability HourlyForecastData = "precipitation_probability"
	UVIndex                  HourlyForecastData = "uv_index"
	UVIndexClearSky          HourlyForecastData = "uv_index_clear_sky"
)

type TemperatureUnit = string

const (
	Fahrenheit TemperatureUnit = "fahrenheit"
	Celcius    TemperatureUnit = "celcius"
)

type DailyForecastQueryParams struct {
	Latitude        string              `url:"latitude"`
	Longitude       string              `url:"longitude"`
	DailyParams     []DailyForecastData `url:"daily"`
	Timezone        string              `url:"timezone"`
	TemperatureUnit TemperatureUnit     `url:"temperature_unit"`
}

type HourlyForecastQueryParams struct {
	Latitude        string               `url:"latitude"`
	Longitude       string               `url:"longitude"`
	HourlyParams    []HourlyForecastData `url:"hourly"`
	Timezone        string               `url:"timezone"`
	TemperatureUnit TemperatureUnit      `url:"temperature_unit"`
}

var ALL_DAILY_DATA = []DailyForecastData{
	MaxTempurature,
	MinTempurature,
}

var ALL_HOURLY_DATA = []HourlyForecastData{
	Tempurature,
	ApparentTempurature,
	RelativeHumidity,
	WeatherCode,
	Precipitation,
	PrecipitationProbability,
	UVIndex,
	UVIndexClearSky,
}

const forecastUrl = "https://api.open-meteo.com/v1/forecast"

func GetDailyForecast(location LocationQuery) (*DailyForecastResponse, error) {
	opt := DailyForecastQueryParams{
		Latitude:        location.Latitude,
		Longitude:       location.Longitude,
		DailyParams:     ALL_DAILY_DATA,
		Timezone:        location.Timezone,
		TemperatureUnit: Fahrenheit,
	}

	q, err := query.Values(opt)

	if err != nil {
		return nil, fmt.Errorf("failed to assemble query: %v", err)
	}

	endpoint := fmt.Sprintf("%s?%s", forecastUrl, q.Encode())

	r, err := http.Get(endpoint)

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching weather: %s", r.Status)
	} else if err != nil {
		return nil, fmt.Errorf("error fetching weather: %v", err)
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	var weather DailyForecastResponse

	if err := json.NewDecoder(r.Body).Decode(&weather); err != nil {
		fmt.Printf("Error parsing response: %s\n", err)
	}

	return &weather, nil
}

func GetHourlyForecast(location LocationQuery) (*HourlyForecastResponse, error) {
	opt := HourlyForecastQueryParams{
		Latitude:        location.Latitude,
		Longitude:       location.Longitude,
		HourlyParams:    ALL_HOURLY_DATA,
		Timezone:        location.Timezone,
		TemperatureUnit: Fahrenheit,
	}

	q, err := query.Values(opt)

	if err != nil {
		return nil, fmt.Errorf("failed to assemble query: %v", err)
	}

	endpoint := fmt.Sprintf("%s?%s", forecastUrl, q.Encode())

	slog.Info(endpoint)

	r, err := http.Get(endpoint)

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching weather: %s", r.Status)
	} else if err != nil {
		return nil, fmt.Errorf("error fetching weather: %v", err)
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	var weather HourlyForecastResponse

	if err := json.NewDecoder(r.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("error parsing response: %s", err)
	}

	return &weather, nil
}
