package main

import (
	"log/slog"
	"net/http"
	"proctorinc/weather-app/internal/api"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func formatDate(value string) string {
	dateLayout := "Monday, Jan 02"
	date, err := time.Parse("2006-01-02", value)

	if err != nil {
		return ""
	}

	return date.Format(dateLayout)
}

func formatTime(value string) string {
	dateLayout := "2006-01-02T15:04"
	timeFormat := "3:04pm"

	datetime, err := time.Parse(dateLayout, value)

	if err != nil {
		return ""
	}

	return datetime.Format(timeFormat)
}

func getLocationQuery(context *gin.Context) api.LocationQuery {
	return api.LocationQuery{
		Name:      context.Query("city"),
		Latitude:  context.Query("lat"),
		Longitude: context.Query("long"),
		Timezone:  context.Query("tz"),
	}
}

func isValidLocationQuery(query api.LocationQuery) bool {
	return query.Name != "" && query.Latitude != "" && query.Longitude != "" && query.Timezone != ""
}

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatDate": formatDate,
		"formatTime": formatTime,
	})
	router.LoadHTMLFiles("templates/location.html", "templates/notfound.html", "templates/weather.html")

	router.GET("/", func(context *gin.Context) {
		query := context.Query("query")

		locations, err := api.SearchLocation(query)

		context.HTML(http.StatusOK, "location.html", gin.H{
			"Query":     query,
			"Locations": locations,
			"Count":     len(locations),
			"Error":     err,
		})
	})

	router.GET("/weather", func(context *gin.Context) {
		query := getLocationQuery(context)

		// Validate query parameters
		if !isValidLocationQuery(query) {
			slog.Error("Invalid location query parameters")
			context.Redirect(http.StatusFound, "/")
			return
		}

		dailyForecast, dailyErr := api.GetDailyForecast(query)
		hourlyForecast, hourlyErr := api.GetHourlyForecast(query)

		context.HTML(http.StatusOK, "weather.html", gin.H{
			"Location":       query,
			"DailyForecast":  dailyForecast,
			"HourlyForecast": hourlyForecast,
			"DailyError":     dailyErr,
			"HourlyError":    hourlyErr,
		})
	})

	router.Run(":8080")
}
