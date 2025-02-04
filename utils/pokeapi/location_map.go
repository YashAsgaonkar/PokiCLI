package pokeapi

import (
	"encoding/json"
	"net/http"
)

//making sure it can be accessed by an instance of client type only
func (c *Client) GetLocations(pageUrl *string) (ResLocationsArea,error){
	url := baseURL+"/location-area"
	if pageUrl != nil{
		url = *pageUrl
	}
	if val,ok:=c.cache.Get(url);ok{
		var locations ResLocationsArea
		if err:=json.Unmarshal(val,&locations);err!=nil{
			return ResLocationsArea{},err
		}
		return locations,nil
	}
	req,err:=http.NewRequest("GET",url,nil)
	if err !=nil{
		return ResLocationsArea{},err
	}
	req.Header.Set("Content-Type","application/json")
	resp,err:=c.httpClient.Do(req)
	if err !=nil{
		return ResLocationsArea{},err
	}
	defer resp.Body.Close()
	var locations ResLocationsArea
	decoder:=json.NewDecoder(resp.Body)
	if err:=decoder.Decode(&locations);err!=nil{
		return ResLocationsArea{},err
	}
	return locations,nil
}