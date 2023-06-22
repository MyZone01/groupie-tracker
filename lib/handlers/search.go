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
	CreationDates []int
}

func Suggestion(res http.ResponseWriter, req *http.Request) {
	searchQuery := strings.ToLower(req.FormValue("search"))

	if searchQuery != "" {
		Names := []string{}
		Locations := []string{}
		Members := []string{}
		FirstAlbums := []string{}
		CreationDates := []int{}

		artists, err := utils.GetArtistList(res)
		if err {
			return
		}

		for _, artist := range artists {
			if strings.Contains(strings.ToLower(artist.Name), searchQuery) {
				Names = append(Names, artist.Name)
			}
			if strings.Contains(strings.ToLower(artist.FirstAlbum), searchQuery) {
				FirstAlbums = append(FirstAlbums, artist.FirstAlbum)
			}
			if fmt.Sprintf("%d", artist.CreationDate) == searchQuery {
				CreationDates = append(CreationDates, artist.CreationDate)
			}
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), searchQuery) {
					Members = append(Members, member)
				}
			}
		}

		data, err := utils.GetLocations(res)
		if err {
			return
		}
		var locations []models.LocationModel
		_err := json.Unmarshal(data, &locations)
		if _err != nil {
			utils.RenderPage("500", nil, res)
			log.Println("❌ Internal Server Error ", err)
			return
		}

		for _, location := range locations {
			for _, _location := range location.Locations {
				__location := strings.Split(_location, "-")
				if len(__location) == 2 {
					city := strings.ToLower(strings.ReplaceAll(__location[0], "_", " "))
					country := strings.ToLower(strings.ReplaceAll(__location[1], "_", " "))
					if strings.Contains(strings.ToLower(city), searchQuery) {
						Locations = append(Names, city)
					}
					if strings.Contains(strings.ToLower(country), searchQuery) {
						Locations = append(Names, country)
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

		log.Println("✅ Suggestion - " + searchQuery)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(jsonResponse)
	}
}

func Search(res http.ResponseWriter, req *http.Request) {

}
