package groupietracker

import (
	"encoding/json"
	models "groupietracker/lib/models"
	utils "groupietracker/lib/utils"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func Map(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/map/*", http.MethodGet) {
		idArtist, err := strconv.Atoi(path.Base(req.URL.Path))
		if err == nil && (idArtist > 0 && idArtist <= 52) {
			pagePath := "map"
			_data, shouldReturn := utils.GetLocations(res)
			if shouldReturn {
				return
			}
			data := models.LocationModel{}
			json.Unmarshal(_data, &data)

			concertLocations := data.Index[idArtist-1].Locations
			var concertCoordinate []models.GPS

			for _, location := range concertLocations {
				city := strings.Split(location, "-")
				city[0] = strings.ReplaceAll(city[0], "_", "+")
				concertCoordinate = append(concertCoordinate, utils.GetCoordinate(city[0]))
			}

			utils.RenderPage(pagePath, &concertCoordinate, res)
			log.Println("✅ Map page get with success")
		} else {
			res.WriteHeader(http.StatusNotFound)
			utils.RenderPage("404", nil, res)
		}
	}
}
