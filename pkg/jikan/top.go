package jikan

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	Last_visible_page int    `json:"last_visible_page"`
	Has_next_page     bool   `json:"has_next_page"`
	Items             []Item `json:"items"`
}

type Response struct {
	Data       []Anime    `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// func GetTopN(n int) (success bool, err error) {
// 	base_url := "https://api.jikan.moe/v4/"

// 	page_to_check := n / 25

// 	return true, nil
// }

func JikanTopAnime(n int) error {
	response, err := http.Get(fmt.Sprintf("%s/top/anime", Base_url))
	if err != nil {
		return err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseObj Response
	json.Unmarshal(responseData, &responseObj)

	for i := 0; i != len(responseObj.Data); i++ {
		fmt.Println(responseObj.Data[i])
	}

	fmt.Println(responseObj.Data[len(responseObj.Data)-1])

	// fmt.Println(string(responseData))
	return nil
}
