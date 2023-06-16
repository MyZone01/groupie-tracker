package groupie_tracker

type RelationModel struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	City    string
	Country string
	Dates   []string
}
