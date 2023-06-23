package groupietracker

import (
	"encoding/json"
	"fmt"
	models "groupietracker/lib/models"
	utils "groupietracker/lib/utils"
	"log"
	"net/http"
	"strings"
)

type SearchSuggestion struct {
	Names         []string
	Members       []string
	Locations     []string
	FirstAlbums   []string
	CreationDates []string
}

func Suggestion(res http.ResponseWriter, req *http.Request) {
	searchQuery := strings.ToLower(req.FormValue("search"))

	if searchQuery != "" {
		Names := []string{}
		Locations := []string{}
		Members := []string{}
		FirstAlbums := []string{}
		CreationDates := []string{}

		artists, err := utils.GetArtistList(res)
		if err {
			return
		}

		log.Println("✅ Suggestion - " + searchQuery)
		for _, artist := range artists {
			if strings.Contains(strings.ToLower(artist.Name), searchQuery) {
				Names = append(Names, fmt.Sprintf("%s@%d", artist.Name, artist.Id))
			}
			if strings.Contains(strings.ToLower(artist.FirstAlbum), searchQuery) {
				FirstAlbums = append(FirstAlbums, fmt.Sprintf("%s@%d", artist.FirstAlbum, artist.Id))
			}
			if fmt.Sprintf("%d", artist.CreationDate) == searchQuery {
				CreationDates = append(CreationDates, fmt.Sprintf("%d@%d", artist.CreationDate, artist.Id))
			}
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), searchQuery) {
					Members = append(Members, fmt.Sprintf("%s@%d", member, artist.Id))
				}
			}
		}

		data, err := utils.GetLocations(res)
		if err {
			return
		}
		var _data models.LocationModel
		_err := json.Unmarshal(data, &_data)
		if _err != nil {
			utils.RenderPage("500", nil, res)
			log.Println("❌ Internal Server Error ", _err)
			return
		}

		locations := _data.Index
		for _, location := range locations {
			for _, _location := range location.Locations {
				__location := strings.Split(_location, "-")
				if len(__location) == 2 {
					city := strings.ToLower(strings.ReplaceAll(__location[0], "_", " "))
					country := strings.ToLower(strings.ReplaceAll(__location[1], "_", " "))
					if strings.Contains(strings.ToLower(city), searchQuery) {
						Locations = append(Locations, fmt.Sprintf("%s@%d", city, location.Id))
					}
					if strings.Contains(strings.ToLower(country), searchQuery) {
						Locations = append(Locations, fmt.Sprintf("%s@%d", country, location.Id))
					}
				}
			}
		}

		response := SearchSuggestion{
			Names,
			Members,
			Locations,
			FirstAlbums,
			CreationDates,
		}

		jsonResponse, _err := json.Marshal(response)
		if _err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(jsonResponse)
	}
}

func Search(res http.ResponseWriter, req *http.Request) {

}
