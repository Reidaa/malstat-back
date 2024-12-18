package jikan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"malstat/scrapper/pkg/utils"
)

type AnimeResponse struct {
	Data Anime `json:"data"`
}

func AnimeByID(id int) (*Anime, error) {
	var responseObj AnimeResponse
	url := fmt.Sprintf("%s/anime/%d", BaseURL, id)

	utils.Debug.Printf("GET %s", url)

	// response, err := http.Get(url)
	response, err := utils.NetClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

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

	// To prevent -> 429 Too Many Requests
	time.Sleep(Cooldown)

	return &responseObj.Data, nil
}
