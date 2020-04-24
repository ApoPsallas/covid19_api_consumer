package app

import (
	"os"

	"io/ioutil"
	"log"
	"net/http"
)

type http_client interface {
	Do(req *http.Request) (*http.Response, error)
}

type Rapidapi_service struct {
	Client http_client
}

//GetAffectedCountries will send a HTTP request
func (api Rapidapi_service) GetAffectedCountriesClient() ([]byte, error) {
	url := os.Getenv("GET_AFFECTED_CITIES_URL")
	apiHost := os.Getenv("API_HOST")
	apiKey := os.Getenv("API_KEY")

	//TODO the above must be injected into the service

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
