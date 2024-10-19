package jikan

type Title struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type Anime struct {
	Mal_id     int     `json:"mal_id"`
	Url        string  `json:"url"`
	Titles     []Title `json:"titles"`
	Type       string  `json:"type"`
	Status     string  `json:"status"`
	Score      float32 `json:"score"`
	ScoredBy   int     `json:"scored_by"`
	Rank       int     `json:"rank"` // Ranking from jikan are unreliable
	Popularity int     `json:"popularity"`
	Members    int     `json:"members"`
	Favorites  int     `json:"favorites"`
}

type Item struct {
	Count    int `json:"count"`
	Total    int `json:"total"`
	Per_page int `json:"per_page"`
}

type Pagination struct {
	Last_visible_page int  `json:"last_visible_page"`
	Has_next_page     bool `json:"has_next_page"`
	Items             Item `json:"items"`
}
