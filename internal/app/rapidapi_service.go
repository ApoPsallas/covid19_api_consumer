package app

import (
	"encoding/json"
	"errors"
	"os"
	"sort"

	"io/ioutil"
	"log"
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RapidapiService struct {
	Client httpClient
}

//GetAffectedCountries will send a HTTP request
func (api RapidapiService) GetAffectedCountries() (*AffectedCountries, error) {
	affectedCountries := AffectedCountries{}
	response, err := api.getAffectedCountriesClient()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &affectedCountries)
	if err != nil {
		return nil, errors.New("Wrong structure of JSON response")
	}
	sort.Strings(affectedCountries.AffectedCountries)
	return &affectedCountries, err
}

func (api RapidapiService) getAffectedCountriesClient() ([]byte, error) {

	url := os.Getenv("GET_AFFECTED_CITIES_URL")
	apiHost := os.Getenv("API_HOST")
	apiKey := os.Getenv("API_KEY")

	req, _ := http.NewRequest(http.MethodGet, "https://"+apiHost+url, nil)

	req.Header.Add("x-rapidapi-host", apiHost)
	req.Header.Add("x-rapidapi-key", apiKey)

	res, err := api.Client.Do(req)
	if err != nil {
		log.Printf("Client Error: %v \n", err.Error())
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Client Error: %v \n", err.Error())
		return nil, err
	}

	return body, nil

}
