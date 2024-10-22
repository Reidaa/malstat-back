package csv

type anime struct {
	Datetime   string  `csv:"datetime"`
	MalID      int     `csv:"mal_id"`
	Title      string  `csv:"title"`
	Type       string  `csv:"type"`
	Rank       int     `csv:"rank"`
	Score      float32 `csv:"score"`
	ScoredBy   int     `csv:"scored_by"`
	Popularity int     `csv:"popularity"`
	Members    int     `csv:"members"`
	Favorites  int     `csv:"favorites"`
}
