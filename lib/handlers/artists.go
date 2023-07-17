package groupietracker

import (
	models "groupietracker/lib/models"
	utils "groupietracker/lib/utils"
	"strconv"

	"log"
	"net/http"
	"path"
)

type ArtistListData struct {
	Artists       []models.ArtistModel
	Locations     []string
	NumberMembers []bool
	CurrentFilter models.FilterParam
}

func ArtistList(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/artists", http.MethodGet) {
		artists, err := utils.GetArtistList(res)
		queryParams := req.URL.Query()
		filters := models.FilterParam{}
		if len(queryParams["locations"]) != 0 {
			location := queryParams.Get("locations")
			filters.Location = location
		}
		if len(queryParams["minFirstAlbum"]) != 0 {
			minFirstAlbum := queryParams.Get("minFirstAlbum")
			nb, err := strconv.Atoi(minFirstAlbum)
			if err != nil {
				return
			}
			filters.MinFirstAlbum = nb
		}
		if len(queryParams["maxFirstAlbum"]) != 0 {
			maxFirstAlbum := queryParams.Get("maxFirstAlbum")
			nb, err := strconv.Atoi(maxFirstAlbum)
			if err != nil {
				return
			}
			filters.MaxFirstAlbum = nb
		}
		if filters.MinFirstAlbum == 1950 && filters.MaxFirstAlbum == 2023 {
			filters.MinFirstAlbum = 0
			filters.MaxFirstAlbum = 0
		}
		if len(queryParams["numberMembers"]) != 0 {
			numberMembers := queryParams["numberMembers"]
			for _, _nb := range numberMembers {
				nb, err := strconv.Atoi(_nb)
				if err != nil {
					return
				}
				filters.NumberMembers = append(filters.NumberMembers, nb)
			}
		}
		if len(queryParams["minCreationDate"]) != 0 {
			creationDate := queryParams.Get("minCreationDate")
			nb, err := strconv.Atoi(creationDate)
			if err != nil {
				return
			}
			filters.MinCreationDate = nb
		}
		if len(queryParams["maxCreationDate"]) != 0 {
			creationDate := queryParams.Get("maxCreationDate")
			nb, err := strconv.Atoi(creationDate)
			if err != nil {
				return
			}
			filters.MaxCreationDate = nb
		}
		if filters.MinCreationDate == 1950 && filters.MaxCreationDate == 2023 {
			filters.MinCreationDate = 0
			filters.MaxCreationDate = 0
		}
		artists = utils.ApplyFilter(artists, filters)
		if err {
			return
		}

		pagePath := "artistsList"
		data := ArtistListData{
			Artists:       artists,
			Locations:     utils.GetAllLocations(),
			NumberMembers: filterMember(filters.NumberMembers),
			CurrentFilter: filters,
		}
		utils.RenderPage(pagePath, &data, res)
		log.Println("✅ All artists get with success")
	}
}

func filterMember(filterMember []int) []bool {
	a := make([]bool, 9)
	for _, v := range filterMember {
		a[v] = true
	}
	return a
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
