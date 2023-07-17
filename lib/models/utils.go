package groupietracker

type FilterParam struct {
	NumberMembers   []int
	Location        string
	MinCreationDate int
	MaxCreationDate int
	MinFirstAlbum   int
	MaxFirstAlbum   int
}

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
	Artist        ArtistModel
}
