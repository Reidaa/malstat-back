package jikan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const Base_url string = "https://api.jikan.moe/v4"

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

type Response struct {
	Data       []Anime    `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func topAnime(page int, animeType string) (*Response, error) {
	var responseObj Response
	var url = fmt.Sprintf("%s/top/anime?page=%d", Base_url, page)

	if animeType != "" {
		url = fmt.Sprintf("%s&type=%s", url, animeType)
	}

	log.Printf("GET %s", url)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 300 {
		return nil, errors.New(response.Status)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseData, &responseObj)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second)
	return &responseObj, nil
}

func TopAnime(n int) (*[]Anime, error) {
	var data []Anime
	types := []string{"tv", "movie", "ova", "tv_special", "special"}

	for t := 0; t != len(types); t++ {
		response, err := topAnime(1, types[t])

		if err != nil {
			return nil, err
		}

		data = append(data, response.Data...)

		for i := 2; i <= n/response.Pagination.Items.Per_page; i++ {
			response, err := topAnime(i, types[t])
			if err != nil {
				return nil, err
			}
			data = append(data, response.Data...)
		}

	}

	for i := 0; i != len(data); i++ {
		fmt.Println(data[i].Titles[0].Title)
	}

	return &data, nil
}
