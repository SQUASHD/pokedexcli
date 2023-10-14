package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (ExtendedAreaData, error) {
	url := baseURL + "/location-area/" + locationName
	if locationName == "" {
		return ExtendedAreaData{}, errors.New("no location name")
	}
	if val, ok := c.cache.Get(locationName); ok {
		locArea := ExtendedAreaData{}
		err := json.Unmarshal(val, &locArea)
		if err != nil {
			return ExtendedAreaData{}, err
		}
		return locArea, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExtendedAreaData{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExtendedAreaData{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExtendedAreaData{}, err
	}

	locArea := ExtendedAreaData{}
	err = json.Unmarshal(dat, &locArea)
	if err != nil {
		return ExtendedAreaData{}, err
	}

	c.cache.Add(locationName, dat)
	return locArea, nil
}
