package groupie_tracker

import (
	"encoding/json"
	models "groupie_tracker/lib/models"
	utils "groupie_tracker/lib/utils"
	"strings"

	"log"
	"net/http"
	"path"
)

func ArtistList(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/artists/", http.MethodGet) {
		url := "https://groupietrackers.herokuapp.com/api/artists"
		data := utils.GetAPI(url)
		var artists []models.ArtistModel
		json.Unmarshal(data, &artists)

		pagePath := "artistsList"
		utils.RenderPage(pagePath, &artists, res)
		log.Println("✅ All artists get with success")
	}
}

func ArtistInfos(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/artist/*", http.MethodGet) {
		idArtist := path.Base(req.URL.Path)

		artistURL := "https://groupietrackers.herokuapp.com/api/artists/" + idArtist
		data1 := utils.GetAPI(artistURL)
		var artist models.ArtistModel
		json.Unmarshal(data1, &artist)

		relationURL := "https://groupietrackers.herokuapp.com/api/relation/" + idArtist
		data := utils.GetAPI(relationURL)
		var _relation models.RelationModel
		json.Unmarshal(data, &_relation)

		locations := []models.Location{}
		for _locations, _dates := range _relation.DatesLocations {
			__locations := strings.Split(_locations, "-")
			var l models.Location
			if len(__locations) == 2 {
				l.City = strings.Title(strings.ReplaceAll(__locations[0], "_", " "))
				l.Country = strings.Title(strings.ReplaceAll(__locations[1], "_", " "))
			}
			l.Dates = _dates
			locations = append(locations, l)
		}

		artist.Relation = locations

		pagePath := "artist"
		utils.RenderPage(pagePath, &artist, res)
		log.Println("✅ Artist " + artist.Name + " infos get with success")
	}
}
