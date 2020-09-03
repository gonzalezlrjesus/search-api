package models

// AppleAPI .
type AppleAPI struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		WrapperType            string  `json:"-"`
		Kind                   string  `json:"kind"`
		ArtistID               int     `json:"-"`
		CollectionID           int     `json:"-"`
		TrackID                int     `json:"-"`
		ArtistName             string  `json:"-"`
		CollectionName         string  `json:"-"`
		TrackName              string  `json:"trackName"`
		CollectionCensoredName string  `json:"-"`
		TrackCensoredName      string  `json:"-"`
		ArtistViewURL          string  `json:"-"`
		CollectionViewURL      string  `json:"-"`
		TrackViewURL           string  `json:"-"`
		PreviewURL             string  `json:"-"`
		ArtworkURL30           string  `json:"-"`
		ArtworkURL60           string  `json:"-"`
		ArtworkURL100          string  `json:"-"`
		CollectionPrice        float64 `json:"-"`
		TrackPrice             float64 `json:"-"`
		ReleaseDate            string  `json:"releaseDate"`
		CollectionExplicitness string  `json:"-"`
		TrackExplicitness      string  `json:"-"`
		DiscCount              int     `json:"-"`
		DiscNumber             int     `json:"-"`
		TrackCount             int     `json:"-"`
		TrackNumber            int     `json:"-"`
		TrackTimeMillis        int     `json:"-"`
		Country                string  `json:"-"`
		Currency               string  `json:"-"`
		PrimaryGenreName       string  `json:"-"`
		IsStreamable           bool    `json:"-"`
		CollectionArtistName   string  `json:"-"`
	} `json:"results"`
}
