package app

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ApoPsallas/covid19_api_consumer/test/mocks"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)

const (
	apiKey = "4p1k3y"
)

func setup() {

	_ = godotenv.Load(Root + "/test/.env.test")
	os.Setenv("GET_AFFECTED_CITIES_URL", "/test")
	os.Setenv("API_HOST", "example.test")
	os.Setenv("API_KEY", apiKey)
}

func TestGetAffectedCountries(t *testing.T) {

	setup()
	expected := NewAffectedCountries([]string{"here", "there"}, "now")
	//response := "{\"affected_countries\":[\"here\",\"there\"],\"statistic_taken_at\":\"now\"}"

	cacheMapper := &mocks.MockCacheMapper{}
	client := &mocks.MockHttpClient{ClientResponse: expected, ApiKey: apiKey}

	api := RapidapiService{Client: client, CacheMapper: cacheMapper}
	actual, err := api.GetAffectedCountries()
	assert.Nil(t, err)
	assert.NotNil(t, actual)

	assert.Equal(t, expected, actual)

}

func TestGetAffectedCountriesWrongApiKey(t *testing.T) {

	setup()
	os.Setenv("API_KEY", "WrongApiKey")

	expected := "Wrong API key"

	cacheMapper := &mocks.MockCacheMapper{}
	client := &mocks.MockHttpClient{}

	api := RapidapiService{Client: client, CacheMapper: cacheMapper}
	actual, err := api.GetAffectedCountries()
	assert.NotNil(t, err)
	assert.Nil(t, actual)
	assert.Equal(t, expected, err.Error())
}

func TestGetAffectedCountriesWrongResponseStructure(t *testing.T) {

	setup()
	response := "Wrong Response"

	expected := "Wrong structure of JSON response"

	cacheMapper := &mocks.MockCacheMapper{}
	client := &mocks.MockHttpClient{ClientResponse: response, ApiKey: apiKey}

	api := RapidapiService{Client: client, CacheMapper: cacheMapper}
	actual, err := api.GetAffectedCountries()
	assert.NotNil(t, err)
	assert.Nil(t, actual)
	assert.Equal(t, expected, err.Error())

}
