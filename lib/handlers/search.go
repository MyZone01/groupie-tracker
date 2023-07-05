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

type SearchResult struct {
	Names         int
	Members       int
	Locations     int
	FirstAlbums   int
	CreationDates int
	Artist        models.ArtistModel
}

func GetSearchSuggestion(res http.ResponseWriter, searchQuery string) (SearchSuggestion, bool) {
	Names := []string{}
	Locations := []string{}
	Members := []string{}
	FirstAlbums := []string{}
	CreationDates := []string{}

	artists, err := utils.GetArtistList(res)
	if err {
		return SearchSuggestion{}, true
	}

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), searchQuery) {
			Names = append(Names, fmt.Sprintf("%s@%d", artist.Name, artist.Id))
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), searchQuery) {
			FirstAlbums = append(FirstAlbums, fmt.Sprintf("%s, %s@%d", artist.Name, artist.FirstAlbum, artist.Id))
		}
		if fmt.Sprintf("%d", artist.CreationDate) == searchQuery {
			CreationDates = append(CreationDates, fmt.Sprintf("%s, %d@%d", artist.Name, artist.CreationDate, artist.Id))
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), searchQuery) {
				Members = append(Members, fmt.Sprintf("%s@%d", member, artist.Id))
			}
		}
	}

	data, err := utils.GetLocations(res)
	if err {
		return SearchSuggestion{}, true
	}
	var _data models.LocationModel
	_err := json.Unmarshal(data, &_data)
	if _err != nil {
		utils.RenderPage("500", nil, res)
		log.Println("❌ Internal Server Error ", _err)
		return SearchSuggestion{}, true
	}

	locations := _data.Index
	for _, location := range locations {
		for _, _location := range location.Locations {
			__location := strings.Split(_location, "-")
			if len(__location) == 2 {
				city := strings.Title(strings.ReplaceAll(__location[0], "_", " "))
				country := strings.Title(strings.ReplaceAll(__location[1], "_", " "))

				if strings.Contains(searchQuery, strings.ToLower(city)) || strings.Contains(strings.ToLower(city), searchQuery) {
					Locations = append(Locations, fmt.Sprintf("%s, %s@%d", city, country, location.Id))
				} else if strings.Contains(searchQuery, strings.ToLower(country)) || strings.Contains(strings.ToLower(country), searchQuery) {
					Locations = append(Locations, fmt.Sprintf("%s, %s@%d", city, country, location.Id))
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
	return response, false
}

func GetSearchResult(res http.ResponseWriter, searchQuery string) (map[string]SearchResult, bool) {
	response := map[string]SearchResult{}

	artists, err := utils.GetArtistList(res)
	if err {
		return map[string]SearchResult{}, true
	}

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), searchQuery) {
			if _, exist := response[artist.Name]; exist {
				response[artist.Name] = SearchResult{
					Names:         response[artist.Name].Names + 1,
					Members:       response[artist.Name].Members,
					Locations:     response[artist.Name].Locations,
					FirstAlbums:   response[artist.Name].FirstAlbums,
					CreationDates: response[artist.Name].CreationDates,
					Artist:        artist,
				}
			} else {
				response[artist.Name] = SearchResult{
					Names:  1,
					Artist: artist,
				}
			}
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), searchQuery) {
			if _, exist := response[artist.Name]; exist {
				response[artist.Name] = SearchResult{
					FirstAlbums:   response[artist.Name].FirstAlbums + 1,
					Names:         response[artist.Name].Names,
					Members:       response[artist.Name].Members,
					Locations:     response[artist.Name].Locations,
					CreationDates: response[artist.Name].CreationDates,
					Artist:        artist,
				}
			} else {
				response[artist.Name] = SearchResult{
					FirstAlbums: 1,
					Artist:      artist,
				}
			}
		}
		if fmt.Sprintf("%d", artist.CreationDate) == searchQuery {
			if _, exist := response[artist.Name]; exist {
				response[artist.Name] = SearchResult{
					CreationDates: response[artist.Name].CreationDates + 1,
					Names:         response[artist.Name].Names,
					Members:       response[artist.Name].Members,
					Locations:     response[artist.Name].Locations,
					FirstAlbums:   response[artist.Name].FirstAlbums,
					Artist:        artist,
				}
			} else {
				response[artist.Name] = SearchResult{
					CreationDates: 1,
					Artist:        artist,
				}
			}
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), searchQuery) {
				if _, exist := response[artist.Name]; exist {
					response[artist.Name] = SearchResult{
						Members:       response[artist.Name].Members + 1,
						Names:         response[artist.Name].Names,
						Locations:     response[artist.Name].Locations,
						FirstAlbums:   response[artist.Name].FirstAlbums,
						CreationDates: response[artist.Name].CreationDates,
						Artist:        artist,
					}
				} else {
					response[artist.Name] = SearchResult{
						Members: 1,
						Artist:  artist,
					}
				}
			}
		}
	}

	data, err := utils.GetLocations(res)
	if err {
		return map[string]SearchResult{}, true
	}
	var _data models.LocationModel
	_err := json.Unmarshal(data, &_data)
	if _err != nil {
		utils.RenderPage("500", nil, res)
		log.Println("❌ Internal Server Error ", _err)
		return map[string]SearchResult{}, true
	}

	locations := _data.Index
	for _, location := range locations {
		for _, _location := range location.Locations {
			__location := strings.Split(_location, "-")
			if len(__location) == 2 {
				city := strings.ToLower(strings.ReplaceAll(__location[0], "_", " "))
				country := strings.ToLower(strings.ReplaceAll(__location[1], "_", " "))
				if strings.Contains(searchQuery, strings.ToLower(city)) || strings.Contains(strings.ToLower(city), searchQuery) {
					artist, _ := utils.GetArtist(location.Id, res)
					if _, exist := response[artist.Name]; exist {
						response[artist.Name] = SearchResult{
							Locations:     response[artist.Name].Locations + 1,
							Names:         response[artist.Name].Names,
							Members:       response[artist.Name].Members,
							FirstAlbums:   response[artist.Name].FirstAlbums,
							CreationDates: response[artist.Name].CreationDates,
							Artist:        artist,
						}
					} else {
						response[artist.Name] = SearchResult{
							Locations: 1,
							Artist:    artist,
						}
					}
				} else if strings.Contains(searchQuery, strings.ToLower(country)) || strings.Contains(strings.ToLower(country), searchQuery) {
					artist, _ := utils.GetArtist(location.Id, res)
					if _, exist := response[artist.Name]; exist {
						response[artist.Name] = SearchResult{
							Locations:     response[artist.Name].Locations + 1,
							Names:         response[artist.Name].Names,
							Members:       response[artist.Name].Members,
							FirstAlbums:   response[artist.Name].FirstAlbums,
							CreationDates: response[artist.Name].CreationDates,
							Artist:        artist,
						}
					} else {
						response[artist.Name] = SearchResult{
							Locations: 1,
							Artist:    artist,
						}
					}
				}
			}
		}
	}

	return response, false
}

func Suggestion(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/suggestion", http.MethodPost) {
		searchQuery := strings.ToLower(req.FormValue("search"))

		if searchQuery != "" {
			response, shouldReturn := GetSearchSuggestion(res, searchQuery)
			if shouldReturn {
				return
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
}

func Search(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/search", http.MethodPost) {
		searchQuery := strings.ToLower(req.FormValue("search"))
		if searchQuery != "" {
			artists, shouldReturn := GetSearchResult(res, searchQuery)
			if shouldReturn {
				return
			}

			if len(artists) == 1 {
				for _, _artist := range artists {
					artist := _artist.Artist
					http.Redirect(res, req, fmt.Sprintf("/artist/%d", artist.Id), http.StatusSeeOther)
				}
			} else {
				pagePath := "search"
				utils.RenderPage(pagePath, &artists, res)
				log.Println("✅ All artists that match the search request")
			}
		}
	}

}
