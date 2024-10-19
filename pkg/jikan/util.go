package jikan

import "time"

const Base_url string = "https://api.jikan.moe/v4"
const Cooldown time.Duration = time.Second

func RemoveUnrankedAnime(in []Anime) (out []Anime) {
	for i := 0; i != len(in); i++ {
		if in[i].Rank != 0 {
			out = append(out, in[i])
		}
	}

	return
}
