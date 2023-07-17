package groupietracker

import (
	"encoding/json"
	"fmt"
	models "groupietracker/lib/models"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	var _artists []models.ArtistModel
	err = json.Unmarshal(data, &_artists)
	if err != nil {
		RenderPage("500", nil, res)
		log.Println("❌ Internal Server Error ", err)
		return nil, true
	}
	for _, artist := range _artists {
		isError, artist := GetLocationFromArtist(artist.Id, res, artist)
		if isError {
			return nil, true
		}
		artists = append(artists, artist)
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

	isError, artist := GetLocationFromArtist(idArtist, res, artist)
	if isError {
		return models.ArtistModel{}, true
	}
	return artist, false
}

func GetLocationFromArtist(idArtist int, res http.ResponseWriter, artist models.ArtistModel) (bool, models.ArtistModel) {
	relationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", idArtist)
	data, err := GetAPI(relationURL)
	if err != nil {
		RenderPage("500", nil, res)
		return true, models.ArtistModel{}
	}

	var _relation models.RelationModel
	err = json.Unmarshal(data, &_relation)
	if err != nil {
		RenderPage("500", nil, res)
		return true, models.ArtistModel{}
	}

	locations := FormatLocations(_relation)

	artist.Relation = locations
	return false, artist
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

func ApplyFilter(artists []models.ArtistModel, filters models.FilterParam) []models.ArtistModel {
	artistsFiltered := []models.ArtistModel{}
	if filters.Location != "" || (filters.MaxFirstAlbum != 0 && filters.MinFirstAlbum != 0) || len(filters.NumberMembers) != 0 || (filters.MinCreationDate != 0 && filters.MaxCreationDate != 0) {
		for _, artist := range artists {
			match := true
			if filters.MaxFirstAlbum != 0 && filters.MinFirstAlbum != 0 {
				_firstAlbum := strings.Split(artist.FirstAlbum, "-")
				if len(_firstAlbum) == 3 {
					firstAlbumYear, err := strconv.Atoi(_firstAlbum[2])
					if err == nil && !(firstAlbumYear >= filters.MinFirstAlbum && firstAlbumYear <= filters.MaxFirstAlbum) {
						match = false
					}
				}
			}
			if filters.MinCreationDate != 0 && filters.MaxCreationDate != 0 {
				if !(artist.CreationDate >= filters.MinCreationDate && artist.CreationDate <= filters.MaxCreationDate) {
					match = false
				}
			}
			if filters.Location != "" {
				locationMatch := false
				_location := strings.Split(filters.Location, "-")
				if len(_location) == 2 {
					city := strings.ReplaceAll(_location[0], "_", " ")
					country := strings.ReplaceAll(_location[1], "_", " ")
					for _, relation := range artist.Relation {
						if strings.Contains(strings.ToLower(relation.City), city) && strings.Contains(strings.ToLower(relation.Country), country) {
							locationMatch = true
							break
						}
					}
				}
				if !locationMatch {
					match = false
				}
			}
			if len(filters.NumberMembers) > 0 {
				validMembers := false
				for _, nb := range filters.NumberMembers {
					if nb == len(artist.Members) {
						validMembers = true
						break
					}
				}
				if !validMembers {
					match = false
				}
			}
			if match {
				artistsFiltered = append(artistsFiltered, artist)
			}
		}
	} else {
		artistsFiltered = artists
	}
	return artistsFiltered
}
