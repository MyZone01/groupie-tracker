package groupietracker

import (
	"encoding/json"
	"fmt"
	models "groupietracker/lib/models"
	"log"
	"net/http"
)

func GetArtistList(res http.ResponseWriter) ([]models.ArtistModel, bool) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	data, err := GetAPI(url)
	if err != nil {
		RenderPage("500", nil, res)
		log.Println("❌ Internal Server Error ", err)
		return nil, true
	}
	var artists []models.ArtistModel
	err = json.Unmarshal(data, &artists)
	if err != nil {
		RenderPage("500", nil, res)
		log.Println("❌ Internal Server Error ", err)
		return nil, true
	}
	return artists, false
}

func GetArtist(idArtist int, res http.ResponseWriter) (models.ArtistModel, bool) {
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", idArtist)
	data1, err := GetAPI(artistURL)
	if err != nil {
		RenderPage("500", nil, nil)
		return models.ArtistModel{}, true
	}

	var artist models.ArtistModel
	err = json.Unmarshal(data1, &artist)
	if err != nil {
		RenderPage("500", nil, res)
		return models.ArtistModel{}, true
	}
	artist.FirstAlbum = FormatDates(artist.FirstAlbum)

	relationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", idArtist)
	data, err := GetAPI(relationURL)
	if err != nil {
		RenderPage("500", nil, res)
		return models.ArtistModel{}, true
	}

	var _relation models.RelationModel
	err = json.Unmarshal(data, &_relation)
	if err != nil {
		RenderPage("500", nil, res)
		return models.ArtistModel{}, true
	}

	locations := FormatLocations(_relation)

	artist.Relation = locations
	return artist, false
}

func GetDates(res http.ResponseWriter) (models.DateModel, bool) {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	data, err := GetAPI(url)
	if err != nil {
		RenderPage("500", nil, res)
		log.Println("❌ Internal Server Error ", err)
		return models.DateModel{}, true
	}
	var _datesList models.DateModel
	err = json.Unmarshal(data, &_datesList)
	if err != nil {
		RenderPage("500", nil, res)
		log.Println("❌ Internal Server Error ", err)
		return models.DateModel{}, true
	}
	return _datesList, false
}

func GetLocations(res http.ResponseWriter) ([]byte, bool) {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	_data, err := GetAPI(url)
	if err != nil {
		RenderPage("500", nil, res)
		log.Println("❌ Internal Server Error ", err)
		return nil, true
	}
	return _data, false
}
