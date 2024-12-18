package jikan

import (
	"encoding/json"
	"fmt"
	"malstat/scrapper/pkg/utils"
)

type AnimeResponse struct {
	Data Anime `json:"data"`
}

func AnimeByID(id int) (*Anime, error) {
	var responseObj AnimeResponse
	url := fmt.Sprintf("%s/anime/%d", BaseURL, id)

	responseData, err := utils.HttpGet(url)
	if err != nil {
		return nil, fmt.Errorf("failed to request %s: %w", url, err)
	}

	err = json.Unmarshal(responseData, &responseObj)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json data: %w", err)
	}

	return &responseObj.Data, nil
}
