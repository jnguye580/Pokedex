package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name

	response, ok := c.cache.Get(url)
	if ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(response, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, dat)

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}
