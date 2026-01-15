package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) CatchPokemon(name string) (Pokemon, error) {
	if name == "" {
		return Pokemon{}, fmt.Errorf("Pokemon name not provided")
	}
	url := baseURL + "/pokemon/" + name
	if cachedPoke, ok := c.cache.Get(url); ok {
		pokeRes := Pokemon{}
		err := json.Unmarshal(cachedPoke, &pokeRes)
		if err != nil {
			return Pokemon{}, fmt.Errorf("failed to unmarshal pokemon data: %w", err)
		}
		return pokeRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("GET request failed: %w", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("GET response not recieved: %w", err)
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to read response body: %w", err)
	}

	if strings.Contains(string(dat), "Not Found") {
		return Pokemon{}, fmt.Errorf("invalid pokemon; (%s) does not exist", name)
	}

	c.cache.Add(url, dat)

	pokeRes := Pokemon{}
	err = json.Unmarshal(dat, &pokeRes)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to unmarshal pokemon data: %w", err)
	}
	return pokeRes, nil
}
