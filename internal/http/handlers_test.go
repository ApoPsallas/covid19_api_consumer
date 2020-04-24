package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ApoPsallas/simpleGoMicroservice/internal/app"
	"github.com/stretchr/testify/assert"
)

type MockApi struct {
	Response []byte
	Err      error
}

func (api MockApi) GetAffectedCountries() ([]byte, error) {
	return api.Response, api.Err
}

func TestAffectedCountriesHandler(t *testing.T) {

	request := httptest.NewRequest("GET", "/affected_countries", nil)
	responseWriter := httptest.NewRecorder()

	affectedCountries := app.NewAffectedCountries([]string{"es", "gr", "it", "uk"}, "now")
	response, _ := json.Marshal(affectedCountries)

	api := MockApi{Response: response, Err: nil}

	handlerStruct := Handlers{Service: api}

	handler := http.HandlerFunc(handlerStruct.AffectedCountriesHandler)

	handler.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, string(response), strings.TrimRight(responseWriter.Body.String(), "\n"))
	//there is a "\n" added from Encoder
}

func TestAffectedCountriesHandlerWrongClientResponseStucture(t *testing.T) {

	request := httptest.NewRequest("GET", "/affected_countries", nil)
	responseWriter := httptest.NewRecorder()

	affectedCountries := "{\"infected_places\":[\"here\",\"there\"],\"when\":\"now\"}"
	response, _ := json.Marshal(affectedCountries)
	expected := "\"Wrong structure of JSON response\""

	api := MockApi{Response: response, Err: nil}

	handlerStruct := Handlers{Service: api}

	handler := http.HandlerFunc(handlerStruct.AffectedCountriesHandler)

	handler.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusInternalServerError, responseWriter.Code)
	assert.Equal(t, expected, strings.TrimRight(responseWriter.Body.String(), "\n"))
	//there is a "\n" added from Encoder
}

func TestAffectedCountriesHandlerServiceUnavailable(t *testing.T) {

	request := httptest.NewRequest("GET", "/affected_countries", nil)
	responseWriter := httptest.NewRecorder()

	affectedCountries := ""
	response, _ := json.Marshal(affectedCountries)
	expected := "\"Service Unavailable\""

	api := MockApi{Response: response, Err: errors.New("Service Unavailable")}

	handlerStruct := Handlers{Service: api}

	handler := http.HandlerFunc(handlerStruct.AffectedCountriesHandler)

	handler.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusServiceUnavailable, responseWriter.Code)
	assert.Equal(t, expected, strings.TrimRight(responseWriter.Body.String(), "\n"))
	//there is a "\n" added from Encoder
}
