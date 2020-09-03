package models

// TvMazeAPI .
type TvMazeAPI []struct {
	Score float64 `json:"score"`
	Show  struct {
		ID           int      `json:"-"`
		URL          string   `json:"-"`
		Name         string   `json:"name"`
		Type         string   `json:"-"`
		Language     string   `json:"-"`
		Genres       []string `json:"-"`
		Status       string   `json:"-"`
		Runtime      int      `json:"-"`
		Premiered    string   `json:"premiered"`
		OfficialSite string   `json:"-"`
		Schedule     struct {
			Time string   `json:"-"`
			Days []string `json:"-"`
		} `json:"-"`
		Rating struct {
			Average float64 `json:"-"`
		} `json:"-"`
		Weight  int `json:"-"`
		Network struct {
			ID      int    `json:"-"`
			Name    string `json:"-"`
			Country struct {
				Name     string `json:"-"`
				Code     string `json:"-"`
				Timezone string `json:"-"`
			} `json:"-"`
		} `json:"-"`
		WebChannel interface{} `json:"-"`
		Externals  struct {
			Tvrage  int    `json:"-"`
			Thetvdb int    `json:"-"`
			Imdb    string `json:"-"`
		} `json:"-"`
		Image struct {
			Medium   string `json:"-"`
			Original string `json:"-"`
		} `json:"-"`
		Summary string `json:"-"`
		Updated int    `json:"-"`
		Links   struct {
			Self struct {
				Href string `json:"-"`
			} `json:"-"`
			Previousepisode struct {
				Href string `json:"-"`
			} `json:"-"`
		} `json:"-"`
	} `json:"show"`
}
