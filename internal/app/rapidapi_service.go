package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/ApoPsallas/covid19_api_consumer/internal/app/repository"
)

const (
	AFFECTED_COUNTRIES = "affected_countries"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RapidapiService struct {
	Client      httpClient
	CacheMapper repository.Cache
}

//GetAffectedCountries will send a HTTP request
func (api RapidapiService) GetAffectedCountries() (*AffectedCountries, error) {
	affectedCountries := AffectedCountries{}
	var err error = nil

	exists, err := api.CacheMapper.Exists(AFFECTED_COUNTRIES)
	if err != nil {
		log.Printf("Redis Error: %v \n", err.Error())
		return nil, err
	}

	if exists == 1 {

		cachedAffectedCountries, err := api.CacheMapper.Get(AFFECTED_COUNTRIES)
		if err != nil {
			log.Printf("Redis Error: %v \n", err.Error())
			return nil, err
		}
		_ = json.Unmarshal([]byte(cachedAffectedCountries), &affectedCountries)

	} else {

		response, err := api.getAffectedCountriesClient()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(response, &affectedCountries)
		if err != nil {
			return nil, errors.New("Wrong structure of JSON response")
		}
		_, err = api.CacheMapper.Set(AFFECTED_COUNTRIES, response, 5*60*time.Second)
		if err != nil {
			log.Printf("Redis Error: %v \n", err.Error())
			return nil, err
		}

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
