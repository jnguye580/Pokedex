package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (ApiResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	response, ok := c.cache.Get(url)
	if ok {
		locationsResp := ApiResponse{}
		err := json.Unmarshal(response, &locationsResp)
		if err != nil {
			return ApiResponse{}, err
		}
		return locationsResp, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return ApiResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ApiResponse{}, err
	}
	c.cache.Add(url, dat)

	locationsResp := ApiResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return ApiResponse{}, err
	}
	return locationsResp, nil
}
