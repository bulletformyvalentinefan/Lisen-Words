package deezer

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func SearchTrack(query string) (*SearchResponse, error) {
	deezerUrl := "https://api.deezer.com/search?q="

	userQuery := url.QueryEscape(query)
	fullUrl := deezerUrl + userQuery

	resp, err := http.Get(fullUrl)
	if err != nil {
    	return nil, err
	}

	defer resp.Body.Close()

	var response SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}