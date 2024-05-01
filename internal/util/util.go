package util

// func ParseLocationFromQuery(context gin.Context) (api.Location, error) {
// 	coordinates, err := parseCoordinates(context.Query("long"), context.Query("lat"))

// 	if err != nil {
// 		return nil, err
// 	}

// 	location := api.Location{
// 		Longitude: coordinates.Longitude,
// 		Latitude: coordinates.Latitude,

// 	}

// 	return
// }

// func parseCoordinates(long string, lat string) (*api.Coordinates, error) {
// 	var latFloat float32
// 	var longFloat float32

// 	latParsed, err := strconv.ParseFloat(lat, 32)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid latitude: %s", lat)
// 	}
// 	latFloat = float32(latParsed)

// 	if err != nil {
// 		return nil, fmt.Errorf("invalid latitude: %s", lat)
// 	}

// 	longParsed, err := strconv.ParseFloat(long, 32)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid longitude: %s", long)
// 	}
// 	longFloat = float32(longParsed)

// 	if err != nil {
// 		return nil, fmt.Errorf("invalid longitude: %s", long)
// 	}

// 	return &api.Coordinates{Latitude: latFloat, Longitude: longFloat}, nil
// }
