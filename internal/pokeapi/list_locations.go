package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (ShortAreaData, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
		locArea := ShortAreaData{}
		err := json.Unmarshal(val, &locArea)
		if err != nil {
			return ShortAreaData{}, err
		}
		return locArea, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShortAreaData{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ShortAreaData{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ShortAreaData{}, err
	}

	locArea := ShortAreaData{}
	err = json.Unmarshal(dat, &locArea)
	if err != nil {
		return ShortAreaData{}, err
	}

	c.cache.Add(url, dat)
	return locArea, nil
}
