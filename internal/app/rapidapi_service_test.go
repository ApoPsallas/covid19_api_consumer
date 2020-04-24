package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"

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

type clientMock struct{}

func (client *clientMock) Do(req *http.Request) (*http.Response, error) {
	reqBody, _ := json.Marshal("{\"Result\":\"Success\"}")
	var err error
	key := req.Header.Get("x-rapidapi-key")
	if key != apiKey {
		reqBody, _ = json.Marshal("")
		err = errors.New("Wrong API key")
	}

	return &http.Response{Body: ioutil.NopCloser(bytes.NewReader(reqBody))}, err
}

func setup() {

	_ = godotenv.Load(Root + "/test/.env.test")
	os.Setenv("GET_AFFECTED_CITIES_URL", "/test")
	os.Setenv("API_HOST", "example.test")
	os.Setenv("API_KEY", apiKey)
}

func TestGetAffectedCountries(t *testing.T) {

	setup()
	expected, _ := json.Marshal("{\"Result\":\"Success\"}")

	client := &clientMock{}

	api := RapidapiService{Client: client}
	actual, err := api.GetAffectedCountries()
	assert.Nil(t, err)
	assert.NotNil(t, actual)

	assert.Equal(t, expected, actual)

}

func TestGetAffectedCountriesWrongApiKey(t *testing.T) {

	setup()
	os.Setenv("API_KEY", "WrongApiKey")

	expected := "Wrong API key"

	client := &clientMock{}

	api := RapidapiService{Client: client}
	actual, err := api.GetAffectedCountries()
	assert.NotNil(t, err)
	assert.Nil(t, actual)
	assert.Equal(t, expected, err.Error())
}
