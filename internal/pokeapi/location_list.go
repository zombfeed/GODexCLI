package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if cachedLoc, ok := c.cache.Get(url); ok {
		fmt.Println("Accessing Cached Locations...")
		locationRes := RespShallowLocations{}
		err := json.Unmarshal(cachedLoc, &locationRes)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("GET request failed: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("GET response not recieved %w", err)
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("failed to read response body: %w", err)
	}
	fmt.Println("Adding to Location Cache...")
	c.cache.Add(url, dat)

	locationRes := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationRes)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("failed to unmarshal location-area data: %w", err)
	}
	return locationRes, nil
}
