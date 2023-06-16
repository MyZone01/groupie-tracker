package groupie_tracker

type RelationModel struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relation struct {
	DatesLocations []Location
}

type Location struct {
	City    string
	Country string
	Dates   []string
}
