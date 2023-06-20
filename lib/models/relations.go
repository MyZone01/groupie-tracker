package groupietracker

type RelationModel struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Event struct {
	City    string
	Country string
	Dates   []string
}
