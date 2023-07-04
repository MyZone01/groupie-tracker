package groupietracker

import (
	utils "groupietracker/lib/utils"
	"strconv"

	"log"
	"net/http"
	"path"
)

func ArtistList(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/artists", http.MethodGet) {
		artists, err := utils.GetArtistList(res)
		if err {
			return
		}

		pagePath := "artistsList"
		utils.RenderPage(pagePath, &artists, res)
		log.Println("✅ All artists get with success")
	}
}

func ArtistInfos(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/artist/*", http.MethodGet) {
		idArtist, err := strconv.Atoi(path.Base(req.URL.Path))
		if err == nil && (idArtist > 0 && idArtist <= 52) {
			artist, err := utils.GetArtist(idArtist, res)
			if err {
				return
			}

			pagePath := "artist"
			utils.RenderPage(pagePath, &artist, res)
			log.Println("✅ Artist " + artist.Name + " infos get with success")
		} else {
			res.WriteHeader(http.StatusNotFound)
			utils.RenderPage("404", nil, res)
		}
	}
}
