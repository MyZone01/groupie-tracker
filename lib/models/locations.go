package groupie_tracker

type LocationInfo struct {
	Ville string
	Info  []ArtistModel
}

type LocationModel struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}
