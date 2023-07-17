package groupietracker

import (
	"encoding/json"
	models "groupietracker/lib/models"
	"strings"
)

func GetCoordinate(locationName string) models.GPS {
	locationCoordinate := models.GPS{}
	url := "https://api.opencagedata.com/geocode/v1/json?q=" + locationName + "&key=b1723ce1cecd403f9139018d0e95c8c5&language=en&pretty=1"

	data, _ := GetAPI(url)

	var GeoData models.DataGeo

	json.Unmarshal(data, &GeoData)

	locationCoordinate.Latitude = GeoData.Results[0].Geometry.Latitude
	locationCoordinate.Longitude = GeoData.Results[0].Geometry.Longitude
	locationCoordinate.Name = strings.Title(strings.ReplaceAll(locationName, "+", " "))
	return locationCoordinate
}
