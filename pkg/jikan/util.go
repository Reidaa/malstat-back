package jikan

import "time"

const (
	BaseURL             string        = "https://api.jikan.moe/v4"
	Cooldown            time.Duration = time.Second
	MaxAllowedHitPerDay int           = 60 * 60 * 24
	MaxSafeHitPerDay    int           = 60 * 60 * 20
)

func RemoveUnrankedAnime(in []Anime) []Anime {
	var out []Anime

	for i := 0; i != len(in); i++ {
		if in[i].Rank != 0 {
			out = append(out, in[i])
		}
	}

	return out
}
