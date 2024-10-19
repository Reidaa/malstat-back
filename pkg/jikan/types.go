package jikan

type Title struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type Anime struct {
	Mal_id  int     `json:"mal_id"`
	Score   float32 `json:"score"`
	Rank    int     `json:"rank"`
	Members int     `json:"members"`
	Titles  []Title `json:"titles"`
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
