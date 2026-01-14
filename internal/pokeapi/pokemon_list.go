package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(location string) (RespShallowEncounters, error) {
	url := baseURL + "/location-area" + location

	if location == "" {
		return RespShallowEncounters{}, fmt.Errorf("location-area not provided")
	}

	if cachedEnc, ok := c.cache.Get(url); ok {
		fmt.Println("Accessing Cached Encounters...")
		encounterRes := RespShallowEncounters{}
		err := json.Unmarshal(cachedEnc, &encounterRes)
		if err != nil {
			return RespShallowEncounters{}, fmt.Errorf("failed to unmarshal pokemon-encounter data: %w", err)
		}
		return RespShallowEncounters{}, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowEncounters{}, fmt.Errorf("GET request failed: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowEncounters{}, fmt.Errorf("GET response not recieved: %w", err)
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowEncounters{}, fmt.Errorf("failed to read response body: %w", err)
	}

	c.cache.Add(url, dat)

	encounterRes := RespShallowEncounters{}
	err = json.Unmarshal(dat, &encounterRes)
	if err != nil {
		return RespShallowEncounters{}, fmt.Errorf("failed to unmarshal pokemon-encounter data: %w")
	}
	return encounterRes, nil
}
