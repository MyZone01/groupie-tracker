package groupie_tracker

import (
	"encoding/json"
	models "groupie_tracker/lib/models"
	utils "groupie_tracker/lib/utils"

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
		artist.FirstAlbum = utils.FormatDates(artist.FirstAlbum)

		relationURL := "https://groupietrackers.herokuapp.com/api/relation/" + idArtist
		data := utils.GetAPI(relationURL)
		var _relation models.RelationModel
		json.Unmarshal(data, &_relation)

		locations := utils.FormatLocations(_relation)

		artist.Relation = locations

		pagePath := "artist"
		utils.RenderPage(pagePath, &artist, res)
		log.Println("✅ Artist " + artist.Name + " infos get with success")
	}
}
