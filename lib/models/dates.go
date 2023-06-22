package groupietracker

type DateModel struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type DateFormat struct {
	Id    string
	Day   string
	Month string
	Year  string
}