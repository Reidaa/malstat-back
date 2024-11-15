package jikan

import "time"

const Base_url string = "https://api.jikan.moe/v4"
const Cooldown time.Duration = time.Second
const MaxAllowedHitPerDay int = 60 * 60 * 24
const MaxSafeHitPerDay int = 60 * 60 * 20

func RemoveUnrankedAnime(in []Anime) []Anime {
	var out []Anime

	for i := 0; i != len(in); i++ {
		if in[i].Rank != 0 {
			out = append(out, in[i])
		}
	}

	return out
}
