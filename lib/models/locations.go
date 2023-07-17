package groupietracker

type LocationModel struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type GPS struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type DataGeo struct {
	Results []struct {
		Geometry struct {
			Latitude  float64 `json:"lat"`
			Longitude float64 `json:"lng"`
		} `json:"geometry"`
	} `json:"results"`
}
