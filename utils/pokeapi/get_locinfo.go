package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocationInfo -
func (c *Client) GetLocationInfo(location string) (ResLocationInfo, error) {
	url:= baseURL + "/location-area/" + location
	if val,ok := c.cache.Get(url); ok {
		var locationInfo ResLocationInfo
		if err := json.Unmarshal(val, &locationInfo); err != nil {
			return ResLocationInfo{}, err
		}
		return locationInfo, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResLocationInfo{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResLocationInfo{}, fmt.Errorf("error in getting location info: %w", err)
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResLocationInfo{}, err
	}
	var locationInfo ResLocationInfo
	if err := json.Unmarshal(dat,&locationInfo); err != nil {
		return ResLocationInfo{}, err
	}
	c.cache.Add(url, dat)
	return locationInfo, nil 
}