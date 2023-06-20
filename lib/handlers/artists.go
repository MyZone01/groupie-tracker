package groupietracker

import (
	"encoding/json"
	models "groupietracker/lib/models"
	utils "groupietracker/lib/utils"

	"log"
	"net/http"
	"path"
)

func ArtistList(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/artists/", http.MethodGet) {
		url := "https://groupietrackers.herokuapp.com/api/artists"
		data, err := utils.GetAPI(url)
		if err != nil {
			utils.RenderPage("500", nil, res)
			log.Println("❌ Internal Server Error ", err)
			return
		}
		var artists []models.ArtistModel
		err = json.Unmarshal(data, &artists)
		if err != nil {
			utils.RenderPage("500", nil, res)
			log.Println("❌ Internal Server Error ", err)
			return
		}

		pagePath := "artistsList"
		res.WriteHeader(http.StatusOK)
		utils.RenderPage(pagePath, &artists, res)
		log.Println("✅ All artists get with success")
	}
}

func ArtistInfos(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/artist/*", http.MethodGet) {
		idArtist := path.Base(req.URL.Path)

		artistURL := "https://groupietrackers.herokuapp.com/api/artists/" + idArtist
		data1, err := utils.GetAPI(artistURL)
		if err != nil {
			utils.RenderPage("500", nil, nil)
			return
		}

		var artist models.ArtistModel
		err = json.Unmarshal(data1, &artist)
		if err != nil {
			utils.RenderPage("500", nil, res)
			return
		}

		if artist.Id > 0 && artist.Id <= 52 {
			artist.FirstAlbum = utils.FormatDates(artist.FirstAlbum)

			relationURL := "https://groupietrackers.herokuapp.com/api/relation/" + idArtist
			data, err := utils.GetAPI(relationURL)
			if err != nil {
				utils.RenderPage("500", nil, res)
				return
			}

			var _relation models.RelationModel
			err = json.Unmarshal(data, &_relation)
			if err != nil {
				utils.RenderPage("500", nil, res)
				return
			}

			locations := utils.FormatLocations(_relation)

			artist.Relation = locations

			pagePath := "artist"
			res.WriteHeader(http.StatusOK)
			utils.RenderPage(pagePath, &artist, res)
			log.Println("✅ Artist " + artist.Name + " infos get with success")
		} else {
			res.WriteHeader(http.StatusNotFound)
			utils.RenderPage("404", nil, res)
		}
	}
}

