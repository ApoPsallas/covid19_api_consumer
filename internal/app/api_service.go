package app

import (
	"os"
	"path/filepath"
	"runtime"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)

type API struct {
	Client *http.Client
}

//GetAffectedCountries will send a HTTP request
func (api *API) GetAffectedCountriesClient() ([]byte, error) {
	err := godotenv.Load(Root + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("GET_AFFECTED_CITIES_URL")
	api_host := os.Getenv("API_HOST")
	api_key := os.Getenv("API_KEY")

	req, _ := http.NewRequest(http.MethodGet, "https://"+api_host+url, nil)

	req.Header.Add("x-rapidapi-host", api_host)
	req.Header.Add("x-rapidapi-key", api_key)

	res, err := api.Client.Do(req)
	if err != nil {
		log.Printf("Client Error: %v \n", err)
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
