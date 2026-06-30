package pokeapi

import (
    "encoding/json"
    "io"
)

func (c *Client) ListLocations(pageURL *string) (ApiResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil{
		url = *pageURL
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

	locationsResp := ApiResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
    	return ApiResponse{}, err
	}
	return locationsResp, nil
}